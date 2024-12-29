package kafka

import (
	"context"
	"errors"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

var _ ISubscriber = (*subscriber)(nil)

type subscriber struct {
	consumer sarama.ConsumerGroup
	topics   []string
	group    string
}

func NewSubscriber(cfg Config) ISubscriber {
	config := sarama.NewConfig()
	config.Version, _ = sarama.ParseKafkaVersion(cfg.Version)
	config.ClientID = cfg.GroupId

	config.Consumer.Group.Heartbeat.Interval = 5 * time.Second
	config.Consumer.Group.Session.Timeout = 15 * time.Second
	config.Consumer.MaxProcessingTime = 300 * time.Millisecond
	config.Consumer.Return.Errors = true

	if cfg.Consumer.GroupHeartbeatInterval > 0 {
		config.Consumer.Group.Heartbeat.Interval = time.Duration(cfg.Consumer.GroupHeartbeatInterval) * time.Second
	}
	if cfg.Consumer.GroupSessionTimeout > 0 {
		config.Consumer.Group.Session.Timeout = time.Duration(cfg.Consumer.GroupSessionTimeout) * time.Second
	}
	if cfg.Consumer.MaxProcessingTime > 0 {
		config.Consumer.MaxProcessingTime = time.Duration(cfg.Consumer.MaxProcessingTime) * time.Millisecond
	}
	if cfg.Consumer.ReturnErrors != nil {
		config.Consumer.Return.Errors = *cfg.Consumer.ReturnErrors
	}

	if cfg.Acl.Enable {
		config.Net.SASL.Enable = true
		config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA256
		config.Net.SASL.User = cfg.Acl.User
		config.Net.SASL.Password = cfg.Acl.Password
		config.Net.SASL.Handshake = true
		config.Net.TLS.Enable = false
		config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
			return &scramClient{
				HashGeneratorFcn: SHA256,
			}
		}
		config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA256
	}

	if cfg.Newest {
		config.Consumer.Offsets.Initial = sarama.OffsetNewest
	} else {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	//start consumer group
	consumer, err := sarama.NewConsumerGroup(cfg.Addrs, cfg.Group, config)
	if err != nil {
		zap.S().Errorf("init consumer group fail err: %v", err)
		panic(err)
	}

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			if isTemporaryNetworkError(err) {
				zap.S().Warnf("consume error: %v", err)
			} else {
				zap.S().Errorf("consume error: %v", err)
			}
		}
	}()

	return &subscriber{
		consumer: consumer,
		topics:   cfg.Topics,
		group:    cfg.Group,
	}
}

func (sub *subscriber) Read(callback func(context.Context, string, []byte) error) {
	if sub.consumer == nil {
		zap.S().Error("consumer nil -> missing consumer")
		return
	}
	ctx, cancel := context.WithCancel(context.Background())

	handler := &consumerHandler{
		fc:    callback,
		ready: make(chan bool),
		group: sub.group,
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := sub.consumer.Consume(ctx, sub.topics, handler); err != nil {
				zap.S().Errorf("kafka consume topics err: %v", err)
				return
			}
			if ctx.Err() != nil {
				zap.S().Info(ctx.Err())
				return
			}
			handler.ready = make(chan bool)
		}
	}()

	<-handler.ready
	zap.S().Debug("kafka consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		zap.S().Info("terminating: context cancelled - ", sub.group)
	case <-sigterm:
		zap.S().Info("terminating: via signal - ", sub.group)
		sub.consumer.PauseAll()
	}

	cancel()
	wg.Wait()
	sub.Close()
}

func (s subscriber) Close() {
	if err := s.consumer.Close(); err != nil {
		zap.S().Errorf("Error closing client: %v", err)
	}
}

func isTemporaryNetworkError(err error) bool {
	return errors.Is(err, io.ErrUnexpectedEOF) ||
		errors.Is(err, syscall.ECONNREFUSED) ||
		errors.Is(err, syscall.ECONNRESET) ||
		errors.Is(err, syscall.EPIPE)
}

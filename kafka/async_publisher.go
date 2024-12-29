package kafka

import (
	"encoding/json"
	"time"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)


var _ IPublisher = (*asyncPublisher)(nil)

type asyncPublisher struct {
	producer        sarama.AsyncProducer
	maxMessageBytes int64
}

func NewAsyncPublisher(cfg Config) IPublisher {
	log := zap.L().With(zap.String("prefix", "NewAsyncPublisher")).Sugar()

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Flush.Frequency = 50 * time.Millisecond
	if cfg.MaxMessageBytes > 0 {
		config.Producer.MaxMessageBytes = int(cfg.MaxMessageBytes * MB)
	}
	if cfg.Compress {
		config.Producer.Compression = sarama.CompressionGZIP
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
	config.Metadata.Full = false

	producer, err := sarama.NewAsyncProducer(cfg.Addrs, config)
	if err != nil {
		log.Errorf("NewAsyncProducer fail: %v", err)
		return nil
	}

	go func() {
		for err := range producer.Errors() {
			log.Errorf("Failed to write entry err: %v", err)
		}
	}()

	return &asyncPublisher{
		producer:        producer,
		maxMessageBytes: int64(config.Producer.MaxMessageBytes),
	}
}

func (p *asyncPublisher) Write(v interface{}, topic string) error {
	log := zap.L().With(zap.String("prefix", "Write")).Sugar()

	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	p.producer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}
	if int64(len(data)) > warningMessageSize || int64(len(data)) >= p.maxMessageBytes {
		log.Warnf("topic: %v, messageLen=%v, maxMessageByte=%v", topic, len(data), p.maxMessageBytes)
	}
	return nil
}

func (p *asyncPublisher) WriteWithKey(v interface{}, key, topic string) error {
	log := zap.L().With(zap.String("prefix", "WriteWithKey")).Sugar()

	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	p.producer.Input() <- &sarama.ProducerMessage{
		Key:   sarama.StringEncoder(key),
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}
	if int64(len(data)) > warningMessageSize || int64(len(data)) >= p.maxMessageBytes {
		log.Warnf("topic: %v, messageLen=%v, maxMessageByte=%v", topic, len(data), p.maxMessageBytes)
	}
	return nil
}

func (p *asyncPublisher) WriteStringWithKey(v, key, topic string) error {
	log := zap.L().With(zap.String("prefix", "WriteStringWithKey")).Sugar()

	p.producer.Input() <- &sarama.ProducerMessage{
		Key:   sarama.StringEncoder(key),
		Topic: topic,
		Value: sarama.StringEncoder(v),
	}
	if int64(len(v)) > warningMessageSize || int64(len(v)) >= p.maxMessageBytes {
		log.Warnf("topic: %v, messageLen=%v, maxMessageByte=%v", topic, len(v), p.maxMessageBytes)
	}

	return nil
}

func (p *asyncPublisher) Close() {
	p.producer.AsyncClose()
}

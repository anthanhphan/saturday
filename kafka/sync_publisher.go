package kafka

import (
	"encoding/json"
	"time"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

var _ IPublisher = (*syncPublisher)(nil)

type syncPublisher struct {
	producer        sarama.SyncProducer
	maxMessageBytes int64
}

func NewSyncPublisher(cfg Config) IPublisher {
	log := zap.L().With(zap.String("prefix", "NewSyncPublisher")).Sugar()

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Flush.Frequency = 50 * time.Millisecond
	config.Producer.Return.Successes = true

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

	producer, err := sarama.NewSyncProducer(cfg.Addrs, config)
	if err != nil {
		log.Errorf("NewSyncProducer fail: %v", err)
		return nil
	}

	return &syncPublisher{
		producer:        producer,
		maxMessageBytes: int64(config.Producer.MaxMessageBytes),
	}
}

func (p *syncPublisher) Write(v interface{}, topic string) error {
	log := zap.L().With(zap.String("prefix", "Write")).Sugar()

	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if _, _, err := p.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}); err != nil {
		log.Errorf("SendMessage fail: %v", err)
		return err
	}
	return nil
}

func (p *syncPublisher) WriteWithKey(v interface{}, key, topic string) error {
	log := zap.L().With(zap.String("prefix", "WriteWithKey")).Sugar()

	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if _, _, err := p.producer.SendMessage(&sarama.ProducerMessage{
		Key:   sarama.StringEncoder(key),
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}); err != nil {
		log.Errorf("SendMessage fail: %v", err)
		return err
	}
	return nil
}

func (p *syncPublisher) WriteStringWithKey(v, key, topic string) error {
	log := zap.L().With(zap.String("prefix", "WriteStringWithKey")).Sugar()

	if _, _, err := p.producer.SendMessage(&sarama.ProducerMessage{
		Key:   sarama.StringEncoder(key),
		Topic: topic,
		Value: sarama.StringEncoder(v),
	}); err != nil {
		log.Errorf("SendMessage fail: %v", err)
		return err
	}
	return nil
}

func (p *syncPublisher) Close() {
	p.producer.Close()
}

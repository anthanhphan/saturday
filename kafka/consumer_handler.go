package kafka

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type consumerHandler struct {
	fc    func(context.Context, string, []byte) error
	ready chan bool
	group string
}

func (c *consumerHandler) Setup(sarama.ConsumerGroupSession) error {
	zap.S().Info("setup consumer group handler - ", c.group)
	close(c.ready)
	return nil
}
func (c *consumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	zap.S().Info("cleanup consumer group handler - ", c.group)
	return nil
}

func (c *consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg, ok := <-claim.Messages():
			if ok {
				requestID := uuid.New().String()
				ctx := context.Background()
				ctx = context.WithValue(ctx, "request_id", requestID)
				err := c.fc(ctx, msg.Topic, msg.Value)
				if err != nil {
					continue
				}
				session.MarkMessage(msg, "")
			}
		case <-session.Context().Done():
			return nil
		}
	}
}

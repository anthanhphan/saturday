package main

import (
	"context"
	"fmt"

	"github.com/anthanhphan/saturday/kafka"
	"github.com/anthanhphan/saturday/logger"
)

func handler(ctx context.Context, topic string, value []byte) error {
	fmt.Println(string(value))
	return nil
}

func main() {
	logInstance, undo := logger.InitLogger(&logger.Config{
		DisableCaller:     false,
		DisableStacktrace: true,
		EnableDevMode:     true,
		Level:             logger.LevelInfo,
		Encoding:          logger.EncodingConsole,
	})
	defer func() {
		_ = logInstance.Sync()
	}()
	defer undo()

	subscriber := kafka.NewSubscriber(kafka.Config{
		Addrs:  []string{"localhost:9094"},
		Group:  "test",
		Topics: []string{"test_topic"},
	})

	subscriber.Read(handler)
}

package main

import (
	"github.com/anthanhphan/saturday/kafka"
	"github.com/anthanhphan/saturday/logger"
	"github.com/google/uuid"
)

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

	publisher := kafka.NewSyncPublisher(kafka.Config{
		Addrs: []string{"localhost:9094"},
	})

	publisher.Write(map[string]interface{}{
		"name": "Peter",
		"age":  20,
		"id":   uuid.New().String(),
	}, "test_topic")

}

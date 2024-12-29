package kafka

import "context"

type IPublisher interface {
	Write(value interface{}, topic string) error
	WriteWithKey(value interface{}, key, topic string) error
	WriteStringWithKey(value, key, topic string) error
	Close()
}

type ISubscriber interface {
	Read(callback func(context.Context, string, []byte) error)
	Close()
}

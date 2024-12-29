package kafka

const (
	_        = iota
	KB int64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

const warningMessageSize = 1024 * 1024

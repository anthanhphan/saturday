package kafka

type Config struct {
	Addrs           []string `json:"addrs" yaml:"addrs"`
	Topics          []string `json:"topics" yaml:"topics"`
	Group           string   `json:"group" yaml:"group"`
	GroupId         string   `json:"group_id" yaml:"group_id"`
	MaxMessageBytes int64    `json:"max_message_bytes" yaml:"max_message_bytes"`
	Compress        bool     `json:"compress" yaml:"compress"`
	Newest          bool     `json:"newest" yaml:"newest"`
	Version         string   `json:"version" yaml:"version"`
	Consumer        struct {
		GroupHeartbeatInterval int64 `json:"group_heartbeat_interval" yaml:"group_heartbeat_interval"`
		GroupSessionTimeout    int32 `json:"group_session_timeout" yaml:"group_session_timeout"`
		MaxProcessingTime      int32 `json:"max_processing_time" yaml:"max_processing_time"`
		ReturnErrors           *bool `json:"return_errors" yaml:"return_errors"`
	}
	Acl struct {
		Enable   bool   `json:"enable" yaml:"enable"`
		User     string `json:"user" yaml:"user"`
		Password string `json:"password" yaml:"password"`
	}
}

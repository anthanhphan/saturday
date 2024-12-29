package kafka

type Config struct {
	Addrs           []string
	Topics          []string
	Group           string
	GroupId         string
	MaxMessageBytes int64
	Compress        bool
	Newest          bool
	Version         string
	Consumer        struct {
		GroupHeartbeatInterval int64
		GroupSessionTimeout    int32
		MaxProcessingTime      int32
		ReturnErrors           *bool
	}
	Acl struct {
		Enable   bool
		User     string
		Password string
	}
}

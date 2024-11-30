package models

type Postgres struct {
	Host                  string `yaml:"host"`
	Port                  int64  `yaml:"port"`
	Database              string `yaml:"database"`
	User                  string `yaml:"user"`
	Password              string `yaml:"password"`
	TimeZone              string `yaml:"time_zone"`
	SSLCertPath           string `yaml:"ssl_cert_path"`
	SSLKeyPath            string `yaml:"ssl_key_path"`
	SSLRootCertPath       string `yaml:"ssl_root_cert_path"`
	SSLMode               string `yaml:"ssl_mode"`
	LogLevel              string `yaml:"log_level"`
	SlowQueryThreshold    int64  `yaml:"slow_query_threshold"`
	MaxOpenConnections    int    `yaml:"max_open_connections"`
	MaxIdleConnections    int    `yaml:"max_idle_connections"`
	ConnectionMaxIdleTime int64  `yaml:"connection_max_idle_time"`
	ConnectionMaxLifeTime int64  `yaml:"connection_max_life_time"`
	ConnectionTimeout     int64  `yaml:"connection_timeout"`
}

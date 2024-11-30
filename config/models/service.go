package models

type Service struct {
	Host    string `yaml:"host" json:"host"`
	Port    int64  `yaml:"port" json:"port"`
	Name    string `yaml:"name" json:"name"`
	Timeout int64  `yaml:"timeout" json:"timeout"`
}

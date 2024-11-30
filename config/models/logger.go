package models

type Logger struct {
	DisableCaller     bool   `yaml:"disable_caller" json:"disable_caller"`
	DisableStacktrace bool   `yaml:"disable_stacktrace" json:"disable_stacktrace"`
	EnableDevMode     bool   `yaml:"enable_dev_mode" json:"enable_dev_mode"`
	Level             string `yaml:"level" json:"level"`
	Encoding          string `yaml:"encoding" json:"encoding"`
}

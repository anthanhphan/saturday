package mail

type MailConfig struct {
	MailFrom     string `json:"mail_from" yaml:"mail_from"`
	MailServer   string `json:"mail_server" yaml:"mail_server"`
	MailPort     int64  `json:"mail_port" yaml:"mail_port"`
	MailPassword string `json:"mail_password" yaml:"mail_password"`
}

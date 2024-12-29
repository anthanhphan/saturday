package mail

type IMail interface {
	SendMail(to, subject, templatePath string, data interface{}) error
}

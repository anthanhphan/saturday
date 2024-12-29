package mail

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"
)

type MailConfig struct {
	MailFrom     string `json:"mail_from" yaml:"mail_from"`
	MailServer   string `json:"mail_server" yaml:"mail_server"`
	MailPort     int64  `json:"mail_port" yaml:"mail_port"`
	MailPassword string `json:"mail_password" yaml:"mail_password"`
}

func (m *MailConfig) InitMail() MailConfig {
	return MailConfig{
		MailFrom:     m.MailFrom,
		MailServer:   m.MailServer,
		MailPort:     m.MailPort,
		MailPassword: m.MailPassword,
	}
}

func (m *MailConfig) SendMail(to, subject, templatePath string, data interface{}) error {
	if len(to) == 0 || len(subject) == 0 || len(templatePath) == 0 {
		return errors.New("to, subject, templatePath can not empty")
	}

	body, err := parseTemplate(templatePath, data)
	if err != nil {
		return err
	}

	var messages []string

	messages = append(messages, "From: "+m.MailFrom+"\r")
	messages = append(messages, "To: "+to+"\r")
	messages = append(messages, "Subject: "+subject+"\r")
	messages = append(messages, "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n")
	messages = append(messages, body+"\r")

	msg := []byte(strings.Join(messages, "\n"))
	mailAuth := fmt.Sprintf("%s:%d", m.MailServer, m.MailPort)

	err = smtp.SendMail(mailAuth,
		smtp.PlainAuth("", m.MailFrom, m.MailPassword, m.MailServer), m.MailFrom, []string{to}, msg)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func parseTemplate(templatePath string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

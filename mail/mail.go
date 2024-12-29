package mail

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"
)

func NewMail(mailConfig MailConfig) IMail {
	return MailConfig{
		MailFrom:     mailConfig.MailFrom,
		MailServer:   mailConfig.MailServer,
		MailPort:     mailConfig.MailPort,
		MailPassword: mailConfig.MailPassword,
	}
}

func (m MailConfig) SendMail(to, subject, templatePath string, data interface{}) error {
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

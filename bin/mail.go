package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func sendmail(to, subject, body string) (err error) {
	msg := fmt.Sprintf(`From: %q <%s>
To: %s
Subject: %s
Content-type: text/plain; charset=UTF-8

%s
`, sMailName, sMailFrom, to, subject, body)

	if sSmtpUser != "" {
		auth := smtp.PlainAuth("", sSmtpUser, sSmtpPass, strings.Split(sSmtpServ, ":")[0])
		err = smtp.SendMail(sSmtpServ, auth, sMailFrom, []string{to}, []byte(msg))
	} else {
		err = smtp.SendMail(sSmtpServ, nil, sMailFrom, []string{to}, []byte(msg))
	}
	return
}

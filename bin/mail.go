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
`, mailName, mailFrom, to, subject, body)

	if smtpUser != "" {
		auth := smtp.PlainAuth("", smtpUser, smtpPass, strings.Split(smtpServ, ":")[0])
		err = smtp.SendMail(smtpServ, auth, mailFrom, []string{to}, []byte(msg))
	} else {
		err = smtp.SendMail(smtpServ, nil, mailFrom, []string{to}, []byte(msg))
	}
	return
}

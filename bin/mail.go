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
`, cfg.Mailname, cfg.Mailfrom, to, subject, body)

	if cfg.Smtpuser != "" {
		auth := smtp.PlainAuth("", cfg.Smtpuser, cfg.Smtppass, strings.Split(cfg.Smtpserv, ":")[0])
		err = smtp.SendMail(cfg.Smtpserv, auth, cfg.Mailfrom, []string{to}, []byte(msg))
	} else {
		err = smtp.SendMail(cfg.Smtpserv, nil, cfg.Mailfrom, []string{to}, []byte(msg))
	}
	return
}

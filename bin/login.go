package main

import (
	"github.com/dchest/authcookie"

	"crypto/sha256"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/smtp"
	"net/url"
	"os"
	"strings"
	"time"
)

func loginForm() {
	b, err := ioutil.ReadFile("../templates/login.html")
	if xx(err) {
		return
	}
	headers()
	fmt.Print(string(b))

	// TODO: remove expired log-in attempts

}

func loginRequest() {
	email := strings.TrimSpace(gReq.FormValue("email"))

	if email == "" {
		x(fmt.Errorf("Missing e-mail address"))
		return
	}

	if !strings.Contains(email, "@") {
		x(fmt.Errorf("Invalid e-mail address"))
		return
	}

	tx, err := gDB.Begin()
	if xx(err) {
		return
	}

	ehash := getHash(email)
	name := getName(email)

	auth := rand16()
	sec := rand16()
	expires := time.Now().Add(time.Hour).Format("2006-01-02T15:04:05")

	result, err := gDB.Exec(fmt.Sprintf("UPDATE `users` SET `sec` = %q, `pw` = %q, `expires` = %q WHERE `email` = %q",
		sec, auth, expires, ehash))
	if x(err) {
		tx.Rollback()
		return
	}
	n, err := result.RowsAffected()
	if x(err) {
		tx.Rollback()
		return
	}
	if n == 0 {
		_, err = tx.Exec("INSERT INTO `users` (email, name, sec, pw, expires) VALUES (?,?,?,?,?);",
			ehash, name, sec, auth, expires)
		if xx(err) {
			tx.Rollback()
			return
		}
	}

	xx(tx.Commit())

	err = sendmail(
		email,
		"Log in",
		fmt.Sprintf(
			"Go to this URL to log in: %sbin/?action=login&pw=%s",
			sBaseUrl, url.QueryEscape(auth)))
	if xx(err) {
		return
	}

	b, err := ioutil.ReadFile("../templates/loginmailed.html")
	if xx(err) {
		return
	}

	t, err := template.New("foo").Parse(string(b))
	if xx(err) {
		return
	}
	headers()
	xx(t.Execute(os.Stdout, email))
}

func login() {

	pw := strings.TrimSpace(gReq.FormValue("pw"))
	if pw == "" {
		pw = "none" // anders kan iemand zonder password inloggen
	}

	rows, err := gDB.Query(fmt.Sprintf("SELECT `email`,`name`,`uid`,`sec` FROM `users` WHERE `pw` = %q", pw))
	if x(err) {
		return
	}
	if rows.Next() {
		err := rows.Scan(&gUserHash, &gUserName, &gUserID, &gUserSec)
		rows.Close()
		if x(err) {
			return
		}
		_, err = gDB.Exec(fmt.Sprintf("UPDATE `users` SET `pw` = '', `expires` = '' WHERE `email` = %q", gUserHash))
		if x(err) {
			return
		}
		gUserAuth = true
		gLocation = true
		headers()
	} else {
		x(fmt.Errorf("Log in failed"))
	}
}

func loggedin() {
	if auth, err := gReq.Cookie(sCookiePrefix + "-auth"); err == nil {
		s := strings.SplitN(authcookie.Login(auth.Value, []byte(getRemote()+sSecret)), "|", 2)
		if len(s) == 2 {
			gUserSec = s[0]
			gUserHash = s[1]
			rows, err := gDB.Query(fmt.Sprintf("SELECT `name`,`uid`,`sec` FROM `users` WHERE `email` = %q", gUserHash))
			if err == nil {
				for rows.Next() {
					var sec string
					rows.Scan(&gUserName, &gUserID, &sec)
					if sec == gUserSec {
						gUserAuth = true
					}
				}
			}
		}
	}
}

func rand16() string {
	a := make([]byte, 16)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 16; i++ {
		a[i] = byte(97 + rnd.Intn(26))
	}
	return string(a)
}

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

func getName(s string) string {
	return s[:strings.Index(s, "@")]
}

func getHash(s string) string {
	return fmt.Sprintf("%x", sha256.Sum224([]byte(s)))
}

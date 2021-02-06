package main

import (
	"github.com/dchest/authcookie"

	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
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
}

func loginRequest() {
	email := strings.TrimSpace(req.FormValue("email"))

	if email == "" {
		x(fmt.Errorf("Missing e-mail address"))
		return
	}

	tx, err := db.Begin()
	if xx(err) {
		return
	}

	auth := rand16()
	sec := rand16()
	expires := time.Now().Add(time.Hour).Format("2006-01-02T15:04:05")

	_, err = tx.Exec(fmt.Sprintf("DELETE FROM `users` WHERE `email` = %q", email))
	if xx(err) {
		tx.Rollback()
		return
	}

	_, err = tx.Exec("INSERT INTO `users` (email, sec, pw, expires) VALUES (?,?,?,?);", email, sec, auth, expires)
	if xx(err) {
		tx.Rollback()
		return
	}

	xx(tx.Commit())

	err = sendmail(
		email,
		"Log in",
		fmt.Sprintf(
			"Go to this URL to log in: %sbin/?action=login&pw=%s",
			baseUrl, url.QueryEscape(auth)))
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

	pw := strings.TrimSpace(req.FormValue("pw"))
	if pw == "" {
		pw = "none" // anders kan iemand zonder password inloggen
	}

	rows, err := db.Query(fmt.Sprintf("SELECT `email`,`sec` FROM `users` WHERE `pw` = %q", pw))
	if x(err) {
		return
	}
	if rows.Next() {
		var mail, sec string
		err := rows.Scan(&mail, &sec)
		rows.Close()
		if x(err) {
			return
		}
		_, err = db.Exec(fmt.Sprintf("UPDATE `users` SET `pw` = '', `expires` = '' WHERE `email` = %q", mail))
		if x(err) {
			return
		}
		userAuth = true
		userMail = mail
		userSec = sec
		doLocation = true
		headers()
	} else {
		x(fmt.Errorf("Log in failed"))
	}
}

func loggedin() {
	if auth, err := req.Cookie(cookiePrefix + "-auth"); err == nil {
		s := strings.SplitN(authcookie.Login(auth.Value, []byte(getRemote()+secret)), "|", 2)
		if len(s) == 2 {
			userMail = s[1]
			userSec = s[0]
			rows, err := db.Query(fmt.Sprintf("SELECT `id` FROM `users` WHERE `email` = %q AND `sec` = %q", userMail, userSec))
			if err == nil {
				for rows.Next() {
					userAuth = true
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

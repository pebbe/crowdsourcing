package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
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

func loginRequest(req *http.Request) {
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

func login(req *http.Request) {
	x(fmt.Errorf("Not implemented"))
}

func loggedin() bool {
	return false
}

func rand16() string {
	a := make([]byte, 16)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 16; i++ {
		a[i] = byte(97 + rnd.Intn(26))
	}
	return string(a)
}

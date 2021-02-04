package main

import (
	"fmt"
	"net/http/cgi"
)

func main() {
	fmt.Print(`Content-type: text/html; charset=utf-8
Cache-Control: no-cache
Pragma: no-cache

`)

	md, err := TomlDecodeFile("setup.toml", &cfg)
	if xx(err) {
		return
	}
	if un := md.Undecoded(); len(un) > 0 {
		xx(fmt.Errorf("Error in setup.toml: unknown : %s", un))
		return
	}

	req, err := cgi.Request()
	if xx(err) {
		return
	}
	if xx(req.ParseForm()) {
		return
	}

	if !dbOpen() {
		return
	}
	defer db.Close()

	if req.Method == "GET" {
		// if arg = login: do logindone
		// else if logged in: do userform
		// else: do loginform
		loginForm()
	} else if req.Method == "POST" {
		action := req.FormValue("action")
		if action == "login" {
			loginRequest(req)
		}
		// else if not logged in: do loginform
		// else if action = logout: do logout; do loginform
		// else if action = submit: do process data, do userform
		// else if action = unskip: do resetskips, do userform
		// else error
	} else {
		x(fmt.Errorf("Method not allowed: %s", req.Method))
	}

}

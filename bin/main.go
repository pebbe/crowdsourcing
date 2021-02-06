package main

import (
	"fmt"
	"net/http/cgi"
)

func main() {
	var err error
	req, err = cgi.Request()
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

	loggedin()

	action := req.FormValue("action")
	if req.Method == "GET" {
		if action == "login" {
			login()
		} else if userAuth {
			if action == "logout" {
				doLogout = true
				doLocation = true
				headers()
			} else {
				userForm()
			}
		} else {
			loginForm()
		}
	} else if req.Method == "POST" {
		if action == "login" {
			loginRequest()
		}
		// else if not logged in: do loginform
		// else if action == "logout": do logout; do loginform
		// else if action == "submit": do process data, do userform
		// else if action == "unskip": do resetskips, do userform
		// else error
	} else {
		x(fmt.Errorf("Method not allowed: %s", req.Method))
	}

}

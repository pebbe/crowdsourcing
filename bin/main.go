package main

import (
	"fmt"
	"net/http/cgi"
)

func main() {
	var err error
	gReq, err = cgi.Request()
	if xx(err) {
		return
	}
	if xx(gReq.ParseForm()) {
		return
	}

	if !dbOpen() {
		return
	}
	defer gDB.Close()

	loggedin()

	action := gReq.FormValue("action")
	if gReq.Method == "GET" {
		if action == "login" {
			login()
		} else if gUserAuth {
			if action == "logout" {
				gLogout = true
				gLocation = true
				headers()
			} else {
				userForm()
			}
		} else {
			loginForm()
		}
	} else if gReq.Method == "POST" {
		if action == "login" {
			loginRequest()
		} else if !gUserAuth {
			loginForm()
		} else if action == "submit" {
			submit()
		}
		// else if action == "unskip": do resetskips, do userform
		// else error
	} else {
		x(fmt.Errorf("Method not allowed: %s", gReq.Method))
	}

}

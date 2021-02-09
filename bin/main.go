package main

import (
	"fmt"
	"net/http/cgi"
)

// TODO in all sources:
//   x(error) vs xx(error)
//   better file names and function names

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
		if action == "dl" {
			download()
		} else if action == "login" {
			login()
		} else if gUserAuth {
			userForm()
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
		} else if action == "unskip" {
			unskip()
		}
		// TODO: else error
	} else {
		x(fmt.Errorf("Method not allowed: %s", gReq.Method))
	}

}

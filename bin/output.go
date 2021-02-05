package main

import (
	"github.com/dchest/authcookie"

	"fmt"
	"strings"
	"time"
)

func headers() {
	fmt.Print(`Content-type: text/html; charset=utf-8
Cache-Control: no-cache
Pragma: no-cache
`)

	if doLogout {
		fmt.Printf("Set-Cookie: %s-auth=; expires=Thu, 01 Jan 1970 00:00:00 GMT\n", cookiePrefix)
		userAuth = false
	}

	if userAuth {
		exp := time.Now().AddDate(0, 0, 14).UTC()
		au := authcookie.New(userSec+"|"+userMail, exp, []byte(getRemote()+secret))
		fmt.Printf("Set-Cookie: %s-auth=%s; Expires=%s\n", cookiePrefix, au, exp.Format(time.RFC1123))
	}

	fmt.Println()
}

func getRemote() string {
	a := make([]string, 0, 2)
	if useForwarded {
		if s := req.Header.Get("X-Forwarded-For"); s != "" {
			a = append(a, req.Header.Get("X-Forwarded-For"))
		}
	}
	if useRemote {
		s := req.RemoteAddr
		if i := strings.LastIndex(s, ":"); i > -1 {
			s = s[:i]
		}
		a = append(a, s)
	}
	return strings.Join(a, ", ")
}

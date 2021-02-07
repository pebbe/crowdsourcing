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

	if gLogout {
		fmt.Printf("Set-Cookie: %s-auth=; expires=Thu, 01 Jan 1970 00:00:00 GMT\n", sCookiePrefix)
		gUserAuth = false
	}

	if gUserAuth {
		exp := time.Now().AddDate(0, 0, 14).UTC()
		au := authcookie.New(gUserSec+"|"+gUserMail, exp, []byte(getRemote()+sSecret))
		fmt.Printf("Set-Cookie: %s-auth=%s; Expires=%s\n", sCookiePrefix, au, exp.Format(time.RFC1123))
	}

	fmt.Println()

	if gLocation {
		fmt.Printf(`<head>
<meta http-equiv="Refresh" content="0; URL=%sbin/">
</head>
`,
			sBaseUrl)
	}

}

func getRemote() string {
	a := make([]string, 0, 2)
	if sUseForwarded {
		if s := gReq.Header.Get("X-Forwarded-For"); s != "" {
			a = append(a, gReq.Header.Get("X-Forwarded-For"))
		}
	}
	if sUseRemote {
		s := gReq.RemoteAddr
		if i := strings.LastIndex(s, ":"); i > -1 {
			s = s[:i]
		}
		a = append(a, s)
	}
	return strings.Join(a, ", ")
}

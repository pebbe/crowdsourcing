package main

import (
	"fmt"
)

func headers() {
	fmt.Print(`Content-type: text/html; charset=utf-8
Cache-Control: no-cache
Pragma: no-cache
`)
	// TODO: cookies

	fmt.Println()
}

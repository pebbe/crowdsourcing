package main

import (
	"fmt"
	"io/ioutil"
)

func userForm() {

	b, err := ioutil.ReadFile("../templates/question.html")
	if xx(err) {
		return
	}
	headers()
	fmt.Print(string(b))
}

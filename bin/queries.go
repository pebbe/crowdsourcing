package main

import (
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

func userForm() {

	b, err := ioutil.ReadFile("../templates/question.html")
	if xx(err) {
		return
	}

	t, err := template.New("foo").Parse(string(b))
	if xx(err) {
		return
	}
	headers()
	xx(t.Execute(os.Stdout, questionType{
		UserName: getName(userMail),
		Question: "Is dit een test?",
		Image:    "../img/pic01.jpg",
	}))
}

func getName(s string) string {
	return s[:strings.Index(s, "@")]
}

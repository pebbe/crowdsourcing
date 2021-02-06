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
		Done:     0,
		Skipped:  0,
		Todo:     3,
		ID:       1,
		UserName: getName(userMail),
		Image:    "pic01.jpg",
		Text:     "Mrs Jones",
	}))
}

func getName(s string) string {
	return s[:strings.Index(s, "@")]
}

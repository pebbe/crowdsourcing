package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

func userForm() {

	rows, err := gDB.Query(fmt.Sprintf(`
SELECT * FROM
( SELECT COUNT(*) FROM questions),
( SELECT COUNT(*) FROM answers WHERE uid = %d AND skip = 0),
( SELECT COUNT(*) FROM answers WHERE uid = %d AND skip > 0)`, gUserID, gUserID))
	if xx(err) {
		return
	}
	var total, done, skipped int
	for rows.Next() {
		if xx(rows.Scan(&total, &done, &skipped)) {
			return
		}
	}

	rows, err = gDB.Query(fmt.Sprintf(`
SELECT qid, text, image FROM qc
WHERE qid NOT IN ( SELECT qid FROM answers WHERE uid = %d )
ORDER BY qc._cnt, qid
LIMIT 1`, gUserID))
	if xx(err) {
		return
	}
	var qid int
	var text, image string
	var ok bool
	for rows.Next() {
		if xx(rows.Scan(&qid, &text, &image)) {
			return
		}
		ok = true
	}
	if !ok {
		xx(fmt.Errorf("Missing row"))
	}

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
		Done:     done,
		Skipped:  skipped,
		Todo:     total - done - skipped,
		ID:       qid,
		UserName: getName(gUserMail),
		Image:    image,
		Text:     text,
	}))
}

func getName(s string) string {
	return s[:strings.Index(s, "@")]
}

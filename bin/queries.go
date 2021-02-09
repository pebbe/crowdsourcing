package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
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
	var ok bool
	for rows.Next() {
		if xx(rows.Scan(&total, &done, &skipped)) {
			return
		}
		ok = true
	}
	if !ok {
		xx(fmt.Errorf("Missing row"))
	}

	if done+skipped >= total {
		var src string
		if skipped == 0 {
			src = "finished.html"
		} else {
			src = "skipped.html"
		}
		b, err := ioutil.ReadFile("../templates/" + src)
		if xx(err) {
			return
		}
		t, err := template.New("foo").Parse(string(b))
		if xx(err) {
			return
		}
		headers()
		xx(t.Execute(os.Stdout, questionType{
			Done:    done,
			Skipped: skipped,
			Todo:    total - done - skipped,
		}))
		return
	}

	// CONFIG: text
	// CONFIG: image
	rows, err = gDB.Query(fmt.Sprintf(`
SELECT qid, text, image FROM questions
WHERE qid NOT IN ( SELECT qid FROM answers WHERE uid = %d )
ORDER BY RANDOM()
LIMIT 1`, gUserID))

	if xx(err) {
		return
	}
	var qid int
	// CONFIG: text
	// CONFIG: image
	var text, image string
	for rows.Next() {
		// CONFIG: text
		// CONFIG: image
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
		Done:    done,
		Skipped: skipped,
		Todo:    total - done - skipped,
		Qid:     qid,
		// CONFIG: Text
		// CONFIG: Image
		Text:  text,
		Image: image,
	}))
}

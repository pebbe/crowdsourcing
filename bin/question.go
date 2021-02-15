package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

func question() {

	total, skipped, done, ok := getDone()
	if !ok {
		return
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

	// CONFIG question: image name tagline
	rows, err := gDB.Query(fmt.Sprintf(`
			SELECT qid,
				image,
				name,
				tagline
			FROM questions
			WHERE qid NOT IN ( SELECT qid FROM answers WHERE uid = %d )
			ORDER BY RANDOM()
			LIMIT 1`,
		gUserID))
	if xx(err) {
		return
	}
	var qid int
	// CONFIG question: image name tagline
	var image string
	var name string
	var tagline string
	for rows.Next() {
		// CONFIG question: image name tagline
		if xx(rows.Scan(&qid, &image, &name, &tagline)) {
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

		// CONFIG question: image name tagline
		Image:   image,
		Name:    name,
		Tagline: template.HTML(tagline),
	}))
}

func getDone() (total, skipped, done int, ok bool) {
	if !gUserAuth {
		xx(fmt.Errorf("Not logged in"))
		return
	}
	rows, err := gDB.Query(fmt.Sprintf(`
			SELECT * FROM
			( SELECT COUNT(*) FROM questions),
			( SELECT COUNT(*) FROM answers WHERE uid = %d AND skip > 0),
			( SELECT COUNT(*) FROM answers WHERE uid = %d AND skip = 0)`,
		gUserID, gUserID))
	if xx(err) {
		return
	}
	for rows.Next() {
		if xx(rows.Scan(&total, &skipped, &done)) {
			return
		}
		ok = true
	}
	if !ok {
		xx(fmt.Errorf("Missing row"))
		return
	}
	return
}

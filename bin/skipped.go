package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

func unskip() {

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

	reset := strings.TrimSpace(gReq.FormValue("unskip"))
	if reset == "yes" {
		_, err := gDB.Exec(fmt.Sprintf("DELETE FROM answers WHERE uid = %d AND skip = 1", gUserID))
		if xx(err) {
			return
		}
		userForm()
		return
	}

	if reset == "no" {
		b, err := ioutil.ReadFile("../templates/finished.html")
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
			UserName: getName(gUserMail),
		}))
		return

	}

	userForm()
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func submit() {
	var qid, skip, size int
	var animal, colour string
	var err error
	qid, err = strconv.Atoi(strings.TrimSpace(gReq.FormValue("id")))
	if xx(err) {
		return
	}
	if strings.TrimSpace(gReq.FormValue("skip")) != "" {
		skip = 1
	} else {
		animal = strings.TrimSpace(gReq.FormValue("animal"))
		colour = strings.TrimSpace(gReq.FormValue("colour"))
		size, _ = strconv.Atoi(strings.TrimSpace(gReq.FormValue("size")))
		if animal == "" {
			x(fmt.Errorf("Missing answer for animal"))
			return
		}
		if colour == "" {
			x(fmt.Errorf("Missing answer for colour"))
			return
		}
		if size < 1 || size > 5 {
			x(fmt.Errorf("Missing choice for size"))
			return
		}
	}

	tx, err := gDB.Begin()
	if xx(err) {
		return
	}

	_, err = tx.Exec(fmt.Sprintf("DELETE FROM answers WHERE qid = %d AND uid = %d", qid, gUserID))
	if xx(err) {
		tx.Rollback()
		return

	}

	_, err = tx.Exec("INSERT INTO answers(qid, uid, skip, animal, colour, size) VALUES (?, ?, ?, ?, ?, ?);",
		qid,
		gUserID,
		skip,
		animal,
		colour,
		size)
	if xx(err) {
		tx.Rollback()
		return

	}

	tx.Commit()

	userForm()

}

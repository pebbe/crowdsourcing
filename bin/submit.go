package main

import (
	"fmt"
	"strconv"
	"strings"
)

func submit() {
	// CONFIG: size
	var qid, skip, size int
	// CONFIG: animal
	// CONFIG: colour
	var animal, colour string
	var err error
	qid, err = strconv.Atoi(strings.TrimSpace(gReq.FormValue("qid")))
	if xx(err) {
		return
	}
	if strings.TrimSpace(gReq.FormValue("skip")) != "" {
		skip = 1
	} else {
		// CONFIG: animal
		// CONFIG: colour
		// CONFIG: size
		animal = strings.TrimSpace(gReq.FormValue("animal"))
		colour = strings.TrimSpace(gReq.FormValue("colour"))
		size, _ = strconv.Atoi(strings.TrimSpace(gReq.FormValue("size")))

		// TODO: max string length

		// CONFIG: animal
		if animal == "" {
			x(fmt.Errorf("Missing answer for animal"))
			return
		}

		// CONFIG: colour
		if colour == "" {
			x(fmt.Errorf("Missing answer for colour"))
			return
		}

		// CONFIG: size
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

	// CONFIG: animal
	// CONFIG: colour
	// CONFIG: size
	_, err = tx.Exec("INSERT INTO answers(qid, uid, skip, animal, colour, size) VALUES (?, ?, ?, ?, ?, ?);",
		qid,
		gUserID,
		skip,
		// CONFIG: animal
		// CONFIG: colour
		// CONFIG: size
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

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func submit() {

	// CONFIG: animal colour size
	var animal string
	var colour string
	var size int

	var qid, skip int
	var err error
	qid, err = strconv.Atoi(strings.TrimSpace(gReq.FormValue("qid")))
	if xx(err) {
		return
	}
	if strings.TrimSpace(gReq.FormValue("skip")) != "" {
		skip = 1
	} else {

		// CONFIG: animal colour size

		animal = strings.TrimSpace(gReq.FormValue("animal"))
		colour = strings.TrimSpace(gReq.FormValue("colour"))
		size, _ = strconv.Atoi(strings.TrimSpace(gReq.FormValue("size")))

		if len(animal) > 100 {
			animal = animal[:100]
		}

		if len(colour) > 100 {
			colour = colour[:100]
		}

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

	// CONFIG: animal colour size
	// NOTE: number of question marks must match number of fields and arguments
	_, err = tx.Exec("INSERT INTO answers(qid, uid, skip, animal, colour, size) VALUES (?, ?, ?, ?, ?, ?);",
		qid,
		gUserID,
		skip,
		// CONFIG: animal colour size
		animal,
		colour,
		size)
	if xx(err) {
		tx.Rollback()
		return

	}

	tx.Commit()

	question()

}

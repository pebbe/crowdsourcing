package main

import (
	_ "github.com/mattn/go-sqlite3"

	"database/sql"
)

func dbOpen() bool {
	var err error

	gDB, err = sql.Open("sqlite3", "../db/data.sqlite")
	if xx(err) {
		return false
	}

	return true
}

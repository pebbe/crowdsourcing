package main

import (
	_ "github.com/mattn/go-sqlite3"

	"database/sql"
)

func dbOpen() bool {
	var err error

	db, err = sql.Open("sqlite3", "../db/data.sqlite")
	if xx(err) {
		return false
	}

	// TODO: remove expired log-in attempts

	return true
}

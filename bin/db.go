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

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
	                email   varchar(128) NOT NULL UNIQUE,
	                sec     char(16)     NOT NULL,
	                pw      char(16)     NOT NULL,
	                expires char(20)     NOT NULL);`)
	if xx(err) {
		return false
	}

	// TODO: remove transpired log-in attempts

	return true
}

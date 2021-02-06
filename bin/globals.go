package main

import (
	"database/sql"
	"net/http"
)

var (
	userAuth bool
	userMail string
	userSec  string

	db *sql.DB

	req *http.Request

	doLogout   bool
	doLocation bool
)

package main

import (
	"database/sql"
	"net/http"
)

var (
	gUserAuth bool
	gUserMail string
	gUserSec  string
	gUserID   int

	gDB *sql.DB

	gReq *http.Request

	gLogout   bool
	gLocation bool
)

package main

import ()

type questionType struct {

	// START OF CONFIG OPTIONS

	Image string
	Text  string

	// END OF CONFIG OPTIONS

	ID       int
	UserName string
	Done     int
	Skipped  int
	Todo     int
}

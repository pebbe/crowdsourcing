package main

type questionType struct {

	// CONFIG: Image
	// CONFIG: Text
	Image string
	Text  string

	Qid      int
	UserName string
	Done     int
	Skipped  int
	Todo     int
}

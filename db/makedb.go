package main

import (
	_ "github.com/mattn/go-sqlite3"

	"database/sql"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {

	////////////////////////////////////////////////////////////////
	//
	// Open database
	//

	db, err := sql.Open("sqlite3", "data.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	////////////////////////////////////////////////////////////////
	//
	// Create table: questions
	//

	// CONFIG question: image name tagline
	_, err = db.Exec(`CREATE TABLE questions (
                        qid     INTEGER PRIMARY KEY,
                        label   TEXT UNIQUE,
                        image   TEXT,
                        name    TEXT,
                        tagline TEXT
                      );`)
	if err != nil {
		log.Fatal(err)
	}

	////////////////////////////////////////////////////////////////
	//
	// Read data into table: questions
	//

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	fp, err := os.Open("questions.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(fp)
	r.Comment = '#'
	r.TrimLeadingSpace = true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// CONFIG question: image name tagline
		// NOTE: number of question marks must match number of fields and arguments
		_, err = tx.Exec("INSERT INTO questions(label, image, name, tagline) VALUES (?, ?, ?, ?);",
			record[0], // label (required)
			record[1], // image
			record[2], // name
			record[3]) // tagline
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	////////////////////////////////////////////////////////////////
	//
	// Create table: answers
	//

	// CONFIG answer: animal colour size
	_, err = db.Exec(`CREATE TABLE answers (
                        aid     INTEGER PRIMARY KEY AUTOINCREMENT,
                        qid     INTEGER,
                        uid     INTEGER,
                        skip    INTEGER DEFAULT 0,
                        animal  TEXT DEFAULT "",
                        colour  TEXT DEFAULT "",
                        size    INTEGER DEFAULT 0
                      );`)
	if err != nil {
		log.Fatal(err)
	}

	////////////////////////////////////////////////////////////////
	//
	// Create table: users
	//

	_, err = db.Exec(`CREATE TABLE users (
                        uid     INTEGER PRIMARY KEY AUTOINCREMENT,
                        email   TEXT NOT NULL UNIQUE,
                        sec     TEXT,
                        pw      TEXT,
                        expires TEXT
                      );`)
	if err != nil {
		log.Fatal(err)
	}

	////////////////////////////////////////////////////////////////
	//
	// Make some indexes
	//

	for _, cmd := range []string{
		"CREATE INDEX auid ON answers(uid);",
		"CREATE INDEX askip ON answers(skip);",
		"CREATE UNIQUE INDEX uemail ON users(email);"} {
		_, err = db.Exec(cmd)
		if err != nil {
			log.Fatal(err)
		}
	}

	////////////////////////////////////////////////////////////////
	//
	// Close database
	//

	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}

}

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

	db, err := sql.Open("sqlite3", "data.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	// CONFIG: text
	// CONFIG: image
	_, err = db.Exec(`CREATE TABLE questions (
                        qid   INTEGER PRIMARY KEY,
                        text  TEXT,
                        image TEXT
                      );`)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	fp, err := os.Open("questions.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(fp)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// CONFIG: text + record[1]
		// CONFIG: image + record[2]
		_, err = tx.Exec("INSERT INTO questions(qid, text, image) VALUES (?, ?, ?);", record[0], record[1], record[2])
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	// CONFIG: animal
	// CONFIG: colour
	// CONFIG: size
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

	for _, cmd := range []string{
		"CREATE INDEX auid ON answers(uid);",
		"CREATE INDEX askip ON answers(skip);",
		"CREATE UNIQUE INDEX uemail ON users(email);"} {
		_, err = db.Exec(cmd)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}

}

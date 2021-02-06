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

	// BEGIN CONFIG

	_, err = db.Exec(`CREATE TABLE questions (
                        id    INTEGER PRIMARY KEY,
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
		_, err = tx.Exec("INSERT INTO questions(id, text, image) VALUES (?, ?, ?);", record[0], record[1], record[2])
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: answers

	// END CONFIG

	_, err = db.Exec(`CREATE TABLE users (
                        email   TEXT NOT NULL UNIQUE,
                        sec     TEXT,
                        pw      TEXT,
                        expires TEXT
                      );`)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: indexen

}

package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func download() {
	fmt.Print("Content-type: text/plain; charset=utf-8\n\n")

	// CONFIG: animal colour size
	rows, err := gDB.Query(`
			SELECT qid,
				animal,
				colour,
				size
			FROM answers
			WHERE skip = 0
			ORDER by qid,
				animal,
				colour,
				size
			;`)

	if err != nil {
		fmt.Println(err)
		return
	}

	w := csv.NewWriter(os.Stdout)

	for rows.Next() {
		// CONFIG: animal colour size
		var animal string
		var colour string
		var size int

		var qid int

		// CONFIG: animal colour size
		err := rows.Scan(&qid, &animal, &colour, &size)
		if err != nil {
			rows.Close()
			fmt.Println(err)
			break
		}

		err = w.Write([]string{
			fmt.Sprint(qid), // convert int to string
			// CONFIG: animal colour size
			animal,
			colour,
			fmt.Sprint(size), // convert int to string -- yes, that comma is needed
		})
		if err != nil {
			rows.Close()
			fmt.Println(err)
			break

		}
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}

	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Println(err)
	}

}

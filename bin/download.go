package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func download() {
	fmt.Print("Content-type: text/plain; charset=utf-8\n\n")

	// CONFIG answer: animal colour size
	rows, err := gDB.Query(`
			SELECT label,
				animal,
				colour,
				size
			FROM results
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

	seen := false
	for rows.Next() {
		seen = true

		// CONFIG answer: animal colour size
		var animal string
		var colour string
		var size int

		var label string

		// CONFIG answer: animal colour size
		err := rows.Scan(&label, &animal, &colour, &size)
		if err != nil {
			rows.Close()
			fmt.Println(err)
			break
		}

		err = w.Write([]string{
			label,
			// CONFIG answer: animal colour size
			animal,
			colour,
			fmt.Sprint(size), // Convert int to string
			// Yes, that comma after the last parameter is needed
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

	if !seen {
		fmt.Println("Nothing submitted yet")
	}
}

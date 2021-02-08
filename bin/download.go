package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func download() {
	fmt.Print("Content-type: text/plain; charset=utf-8\n\n")

	// CONFIG: animal
	// CONFIG: colour
	// CONFIG: size
	rows, err := gDB.Query(`
SELECT qid, animal, colour, size
FROM answers
WHERE skip = 0
ORDER by qid, animal, colour, size;`)

	if err != nil {
		fmt.Println(err)
		return
	}

	w := csv.NewWriter(os.Stdout)

	for rows.Next() {
		// CONFIG: animal
		// CONFIG: colour
		// CONFIG: size
		var animal, colour string
		var qid, size int

		err := rows.Scan(&qid, &animal, &colour, &size)
		if err != nil {
			rows.Close()
			fmt.Println(err)
			break
		}

		err = w.Write([]string{
			fmt.Sprint(qid),
			animal,
			colour,
			fmt.Sprint(size)})
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

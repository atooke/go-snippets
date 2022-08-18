package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()
	// read csv values using csv.Reader
	// Iterate through the records
	// Read each record from csv
	r := csv.NewReader(f)
	r.Comma = '|'
	r.FieldsPerRecord = -1
	mat
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if strings.Contains(record[7], "happy") {
			fmt.Println(record)

		}
	}
}

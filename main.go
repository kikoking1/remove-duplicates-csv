package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {

	// FieldName to dupe check
	fieldName := os.Args[1]
	// Pass in the
	absPathOfCSV := filepath.ToSlash(os.Args[2])

	isDuplicate := make(map[string]bool)

	// Typical csv setup vars
	h := make(map[string]int)

	dataRowsHeap := [][]string{}
	firstCsvLine := false
	RequiredCSVColumnHeaders := []string{fieldName}

	csvfile, err := os.Open(absPathOfCSV)
	if err != nil {
		log.Fatal(err.Error())
	}

	r := csv.NewReader(csvfile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if firstCsvLine {

			for x := 0; x < len(record); x++ {
				h[record[x]] = x
			}

			// MAKE SURE ALL DEPENDANT COLUMNS EXIST BEFORE PROCEEDING
			for _, columnName := range RequiredCSVColumnHeaders {
				if _, ok := h[columnName]; !ok {
					panic("\"" + columnName + "\" column header missing from csv.")
				}
			}

			dataRowsHeap = append(dataRowsHeap, record)

		} else if !firstCsvLine {

			// check if the key DOES NOT already exists in the map
			if _, yes := isDuplicate[record[h[fieldName]]]; !yes {
				// put the key in the map.
				isDuplicate[record[h[fieldName]]] = true

				dataRowsHeap = append(dataRowsHeap, record)
			}

		}

		firstCsvLine = false

	}

	csvfile.Close()

	deleteFile(absPathOfCSV)
	writeToResultsCSV(dataRowsHeap, absPathOfCSV, false)

}

func deleteFile(absPathOfCSV string) {
	var errDel = os.Remove(absPathOfCSV)
	if errDel != nil {
		log.Fatal(errDel.Error())
	}
}

func writeToResultsCSV(csvDataRows [][]string, absPathOfCSV string, append bool) {
	writeOptions := os.O_RDWR | os.O_CREATE | os.O_TRUNC
	if append {
		writeOptions = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	}
	file, err := os.OpenFile(absPathOfCSV, writeOptions, 0644)

	if err != nil {
		os.Exit(1)
	}

	defer file.Close()

	w := csv.NewWriter(file)
	w.WriteAll(csvDataRows)
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

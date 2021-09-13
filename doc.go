package main

// This package is a command utility used to remove duplicate rows from csv's based on a single column's values

// Requirements:
/*
This package requires you to pass in 2 arguments via command line to execute properly
Example Usage:
$ go run internal/cmd/remove-duplicates-csv/main.go nameOfColumnToDeDupe project/relative/path/to/csvFile.csv

The process will (in this order):
1. Read all of the rows into memory, filtering out duplicates
2. Delete the old csv
3. Recreate the csv without duplicates in place
*/

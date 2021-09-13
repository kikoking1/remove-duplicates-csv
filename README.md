# Remove Duplicates CSV
This package is a command utility used to remove duplicate rows from csv's based on a single column's values

## Dependency Requirements:
This package requires you to pass in 2 arguments via command line to execute properly
Example Usage: 

`$ go run main.go nameOfColumnToDeDupe /absolute-path/to/csvFile.csv`

## How Exactly Does it Do This?
The process will (in this order):
1. Read all of the rows into memory, filtering out duplicates
2. Delete the old csv
3. Recreate the csv without duplicates from the in memory grid of values, in place of deleted csv


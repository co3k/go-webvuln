package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "You should specify the filepath to an initial sql file and target db file as the argument.\n")
		os.Exit(1)
	}

	sqlpath := os.Args[1]
	data, err := ioutil.ReadFile(sqlpath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "There is an error while reading the specified file:\n %s\n", err)
		os.Exit(1)
	}

	dbpath := os.Args[2]
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "There is an error while opneing the db file:\n %s\n", err)
		os.Exit(1)
	}

	_, err = db.Exec(string(data))
	if err != nil {
		fmt.Fprintf(os.Stderr, "There is an error while processing the specified sql file:\n %s\n", err)
		os.Exit(1)
	}

	fmt.Print("Initialization has been finished")
}

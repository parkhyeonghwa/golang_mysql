package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func dbConnectionAndClose() {
	db, err := sql.Open("mysql", "root:@/golang_database")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	dbConnectionAndClose()
	fmt.Println("SUCCESS!")
}

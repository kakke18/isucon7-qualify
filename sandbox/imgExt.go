package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"log"
)

func main() {
	// connect mysql db
	dsn := "isucon:isucon@tcp(localhost:3306)/isubata?loc=Local&charset=utf8mb4"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %s.", err.Error())
	}
	defer db.Close()

	// get records
	rows, err := db.Query("SELECT `name`, `data` FROM image")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	// output files
	var name string
	var data []byte
	for rows.Next() {
		err2 := rows.Scan(&name, &data)
		if err2 != nil {
			log.Fatalf(err2.Error())
		}
		fmt.Println(name)
		err3 := ioutil.WriteFile("img/"+name, data, 0666)
		if err != nil {
			log.Fatalf(err3.Error())
		}
	}
}

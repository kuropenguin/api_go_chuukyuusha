package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(db-for-go)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println("ping err")
		fmt.Println(err)
	} else {
		fmt.Println("connect to DB")
	}

}

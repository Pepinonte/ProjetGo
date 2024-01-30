package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type Log struct {
	ID     			int64
	date  			string
	module 			string
	fonction 		string
	arguments  	string
	output 			string

}

var db *sql.DB

func Condb() string{
		cfg := mysql.Config{
			User:                 "root",
			Passwd:               "root",
			Net:                  "tcp",
			Addr:                 "localhost:3307",
			DBName:               "projetgo",
		}

		var err error

		db, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			log.Fatal(err)
		}
		pingErr := db.Ping()
		if pingErr != nil {
			log.Fatal(pingErr)
		}
		fmt.Println("Connected!")
		return "Connected!"
}

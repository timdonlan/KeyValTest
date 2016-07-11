package model

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"fmt"
)

var db *sqlx.DB

func OpenDB(dataSourceName string){
	var err error
	db, err = sqlx.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

func ResetDB(dataSourceName string) {

	fmt.Print("%s",dataSourceName)
	os.Remove(dataSourceName)

	var err error
	db, err = sqlx.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	sqlStmt := "Create table KeyVal (key text not null primary key, value text);"
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

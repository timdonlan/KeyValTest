package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

//var db *sqlx.DB

func OpenDB(dataSourceName string) (*KeyValDAL) {

	db, err := sqlx.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	return &KeyValDAL{db}
}

func ResetDB(dataSourceName string) {

	fmt.Print("%s", dataSourceName)
	os.Remove(dataSourceName)


	db, err := sqlx.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

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

func CleanDB(dataSourceName string) {
	os.Remove(dataSourceName)
}

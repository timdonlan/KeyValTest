package model

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var db *sqlx.DB

func ResetDB(dataSourceName string) {
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

func InitDB(dataSourceName string) {
	var err error
	db, err = sqlx.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
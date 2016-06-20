package KeyVal

import (
	"database/sql"
	"log"
	"os"
	"github.com/jmoiron/sqlx"
)

type KeyValData struct{
	Key string
	Value string
}


func Setup(){
	os.Remove("./KeyVal.db")

	db, err := sql.Open("sqlite3", "./KeyVal.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	sqlStmt := "Create table KeyVal (key text not null primary key, value text);"
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}


func Create(key string, value string){
	db, err := sql.Open("sqlite3", "./KeyVal.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into KeyVal(key, value) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(key, value)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func Update(key string, newValue string){
	db, err := sql.Open("sqlite3", "./KeyVal.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("update KeyVal set value = ? where key = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(newValue, key)
	if err != nil {
		log.Fatal(err)
	}

}

func Delete(key string){
	db, err := sql.Open("sqlite3", "./KeyVal.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("delete from KeyVal where key = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(key)
	if err != nil {
		log.Fatal(err)
	}
}

func GetX(key string) KeyValData {

	var retval KeyValData

	//var db *sqlx.DB
	db,err := sqlx.Open("sqlite3", "./KeyVal.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Queryx("select key,value from KeyVal where key = ?", key)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&retval)
	}
	return retval

}

func Get(key string) KeyValData {

	var retval KeyValData

	db, err := sql.Open("sqlite3", "./KeyVal.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select value from KeyVal where key = ?", key)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var value string
		err = rows.Scan(&value)
		if err != nil {
			log.Fatal(err)
		}
		retval.Key = key
		retval.Value = value
	}

	return retval
}

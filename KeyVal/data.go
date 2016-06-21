package KeyVal

import (
	"log"
	"os"
	"github.com/jmoiron/sqlx"
	"fmt"
)

type KeyValData struct{
	Key string
	Value string
}

var db *sqlx.DB

func Setup(){
	os.Remove("./KeyVal.db")
	var err error
	db,err = sqlx.Open("sqlite3", "./KeyVal.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := "Create table KeyVal (key text not null primary key, value text);"
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func Create(key string, value string){
	/*
	db, err = sql.Open("sqlite3", "./KeyVal.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	*/

	stmt, err := db.Prepare("insert into KeyVal(key, value) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(key, value)
	if err != nil {
		log.Fatal(err)
	}
}

func Update(key string, newValue string){
	/*
	db, err := sql.Open("sqlite3", "./KeyVal.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	*/

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
	/*
	db, err := sql.Open("sqlite3", "./KeyVal.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	*/

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

func Get2(key string ) KeyValData{
	var retval KeyValData
	var found bool

	retval, found = Query(retval,"select key,value from KeyVal where key = ?", key).(KeyValData)
	if !found{
		fmt.Print("unable to map data")
	}
	return retval
}

func Get(key string) KeyValData {

	var retval KeyValData

	/*
	db, err := sql.Open("sqlite3", "./KeyVal.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	*/

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

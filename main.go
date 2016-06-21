package main

import ("fmt"
	"os"
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"KeyValTest/model"
	"encoding/json"
)



func main() {

	dbName := "KeyVal.db"

	model.ResetDB(dbName)
	model.InitDB(dbName)
	model.Create("hello", "world")
	model.Create("world2", "hello2")

	model.Update("hello","world3")
	model.Delete("world2")

	dataArray,err := model.GetAll()
	if(err != nil){
		return
	}


	json,err := json.Marshal(dataArray)
	fmt.Printf("%s", json)

	for _, data := range dataArray{
		if(data != nil) {

		fmt.Printf("%s: %s\n",data.Key, data.Value)
		}
	}

	/*
	KeyVal.Setup()
	KeyVal.Create("hello", "world")
	KeyVal.Create("world","hello")

	data := KeyVal.GetX("hello")
	fmt.Printf("Hello: %s\n",data.Value)

	KeyVal.Update("hello","world2")
	data = KeyVal.GetX("hello")
	fmt.Printf("Hello: %s\n",data.Value)

	KeyVal.Delete("hello")

	data = KeyVal.GetX("hello")
	fmt.Printf("Hello: %s\n",data.Value)
	*/

}

func testSQLLite(){
	os.Remove("./foo.db")

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("こんにちわ世界%03d", i))
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()

	rows, err := db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("select name from foo where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	_, err = db.Exec("delete from foo")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	if err != nil {
		log.Fatal(err)
	}

	rows, err = db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
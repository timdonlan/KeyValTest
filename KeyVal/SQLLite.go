package KeyVal

import (

)
import (

)

/*
func Setup(tableQuery string, fileName string){
	os.Remove(fileName)

	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//sqlStmt := "Create table KeyVal (key text not null primary key, value text);"
	_, err = db.Exec(tableQuery)
	if err != nil {
		log.Printf("%q: %s\n", err, tableQuery)
		return
	}
}

func Query(query string, fileName string){
	var retval KeyValData

	db, err := sql.Open("sqlite3", fileName)
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
		retval.Val = value
	}

	return retval
}
*/
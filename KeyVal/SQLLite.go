package KeyVal


import (
	"log"
)

func Query(data interface{}, query string,args...interface{}) interface{}{
	rows, err := db.Queryx(query,args)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&data)
		if err != nil {
			log.Fatal(err)
		}
	}

	return data
}

package model

import "log"

type KeyValData struct{
	Key string
	Value string
}

func GetAll() ([]*KeyValData, error){
	keyValArray := make([]*KeyValData,0)

	rows, err := db.Queryx("select key,value from KeyVal")
	if err != nil {
		return nil,err
	}
	defer rows.Close()

	for rows.Next() {
		tempKeyVal := new(KeyValData)
		err = rows.StructScan(&tempKeyVal)
		if err != nil {
			return nil,err
		}
		keyValArray = append(keyValArray, tempKeyVal)
	}
	return keyValArray,nil
}

func Get(key string) (*KeyValData, error){
	var retval KeyValData

	rows, err := db.Queryx("select key,value from KeyVal where key = ?", key)
	if err != nil {
		return nil,err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&retval)
		if err != nil {
			return nil,err
		}
	}
	return &retval,nil
}

func Create(key string, value string) (*KeyValData, error){
	stmt, err := db.Prepare("insert into KeyVal(key, value) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(key, value)
	if err != nil {
		log.Fatal(err)
	}

	return &KeyValData{key,value},err
}

func Update(key string, newValue string) (*KeyValData, error){
	stmt, err := db.Prepare("update KeyVal set value = ? where key = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(newValue, key)
	if err != nil {
		log.Fatal(err)
	}

	return &KeyValData{key,newValue},err
}

func Delete(key string) (bool, error){
	stmt, err := db.Prepare("delete from KeyVal where key = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(key)
	if err != nil {
		log.Fatal(err)
	}

	return true,err
}

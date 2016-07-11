package model

import (
	"log"
	"errors"
)

type KeyValData struct {
	Key   string
	Value string
}

func GetAll() ([]*KeyValData, error) {
	keyValArray := make([]*KeyValData, 0)

	sqlQuery := `SELECT KEY, VALUE FROM KEYVAL`

	err := db.Select(&keyValArray, sqlQuery)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return keyValArray, nil
}

func GetKeyVal(key string) (*KeyValData, error) {
	var retval KeyValData

	sqlQuery := `SELECT KEY,VALUE FROM KEYVAL WHERE KEY = ?`

	err := db.Get(&retval, sqlQuery, key)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &retval, nil
}

func CreateKeyVal(key string, value string) (*KeyValData, error) {

	sqlStmt := `INSERT INTO KEYVAL (KEY,VALUE) VALUES (?,?)`

	result, err := db.Exec(sqlStmt, key, value)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	if rows,_ := result.RowsAffected(); rows == 0{
		return nil, errors.New("Failed to insert row")
	}

	return &KeyValData{key, value}, nil
}

func UpdateKeyVal(key string, newValue string) (*KeyValData, error) {

	sqlStmt := `UPDATE KEYVAL SET VALUE = ? WHERE KEY = ?`

	result, err := db.Exec(sqlStmt, newValue, key)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	if rows,_ := result.RowsAffected(); rows == 0{
		return nil, errors.New("Failed to update row")
	}

	return &KeyValData{key, newValue}, nil
}

func DeleteKeyVal(key string) (bool, error) {

	sqlStmt := `DELETE FROM KEYVAL WHERE KEY = ?`

	result, err := db.Exec(sqlStmt, key)
	if err != nil {
		log.Print(err)
		return false, err
	}

	if rows,_ := result.RowsAffected(); rows == 0{
		return false, nil
	}

	return true,nil
}

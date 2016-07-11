package model

import "log"

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

	_, err := db.Exec(sqlStmt, key, value) //swallow the result - do something with this?
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &KeyValData{key, value}, nil
}

func UpdateKeyVal(key string, newValue string) (*KeyValData, error) {

	sqlStmt := `UPDATE KEYVAL SET VALUE = ? WHERE KEY = ?`

	_, err := db.Exec(sqlStmt, newValue, key) //swallow the result - do something with this?
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return &KeyValData{key, newValue}, nil
}

func DeleteKeyVal(key string) (bool, error) {

	sqlStmt := `DELETE FROM KEYVAL WHERE KEY = ?`
	_, err := db.Exec(sqlStmt, key) //swallow the result - do something with this?
	if err != nil {
		return false, err
	}
	return true, nil
}

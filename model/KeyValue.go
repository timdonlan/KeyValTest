package model

import (
	"errors"
	"log"
	"github.com/jmoiron/sqlx"
)

type KeyValData struct {
	Key   string
	Value string
}

type KeyValDAL struct{
	db *sqlx.DB
}

type KeyValDALInterface interface{
	GetAll() ([]*KeyValData, error)
	GetKeyVal(key string) (*KeyValData, error)
	CreateKeyVal(key string, value string) (*KeyValData, error)
	UpdateKeyVal(key string, newValue string) (*KeyValData, error)
	DeleteKeyVal(key string) (bool, error)
}

func (t *KeyValDAL) GetAll() ([]*KeyValData, error) {
	keyValArray := make([]*KeyValData, 0)

	sqlQuery := `SELECT KEY, VALUE FROM KEYVAL`

	err := t.db.Select(&keyValArray, sqlQuery)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return keyValArray, nil
}

func (t *KeyValDAL) GetKeyVal(key string) (*KeyValData, error) {
	var retval KeyValData

	sqlQuery := `SELECT KEY,VALUE FROM KEYVAL WHERE KEY = ?`

	err := t.db.Get(&retval, sqlQuery, key)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &retval, nil
}

func (t *KeyValDAL) CreateKeyVal(key string, value string) (*KeyValData, error) {

	sqlStmt := `INSERT INTO KEYVAL (KEY,VALUE) VALUES (?,?)`

	result, err := t.db.Exec(sqlStmt, key, value)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return nil, errors.New("Failed to insert row")
	}

	return &KeyValData{key, value}, nil
}

func (t *KeyValDAL) UpdateKeyVal(key string, newValue string) (*KeyValData, error) {

	sqlStmt := `UPDATE KEYVAL SET VALUE = ? WHERE KEY = ?`

	result, err := t.db.Exec(sqlStmt, newValue, key)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return nil, errors.New("Failed to update row")
	}

	return &KeyValData{key, newValue}, nil
}

func (t *KeyValDAL) DeleteKeyVal(key string) (bool, error) {

	sqlStmt := `DELETE FROM KEYVAL WHERE KEY = ?`

	result, err := t.db.Exec(sqlStmt, key)
	if err != nil {
		log.Print(err)
		return false, err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return false, nil
	}

	return true, nil
}

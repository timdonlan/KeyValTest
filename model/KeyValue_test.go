package model

import (
	"testing"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db,"sqlmock")

	rows := sqlmock.NewRows([]string{"key", "value"}).
	AddRow("hello", "world").
	AddRow("foo", "bar")
	mock.ExpectQuery("SELECT KEY, VALUE FROM KEYVAL").WillReturnRows(rows)

	keyValDAL := &KeyValDAL{sqlxDB}

	keyValData,err := keyValDAL.GetAll()
	if err != nil{
		t.Error(err)
	}

	fmt.Print(keyValData[0].Value)

}

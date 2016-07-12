package model

import "testing"

var testDBName = "default.db"

func initDB() {

	ResetDB(testDBName)
	OpenDB(testDBName)
}

func mockKeyVal(key string, val string) {
	CreateKeyVal(key, val)
}

func TestGetKeyVal(t *testing.T) {
	initDB()

	mockKeyVal("hello", "world")

	keyValData, err := GetKeyVal("hello")
	if err != nil {
		t.Error(err)
	}

	if keyValData.Value != "world" {
		t.Error("Invalid response")
	}

	CleanDB(testDBName)

}

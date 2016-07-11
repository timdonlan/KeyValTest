package main

import (
	"testing"
	"flag"
	"net/http"

	"time"
	"os"
)

func initTestingFlags(){
	sqliteDbName   = flag.String("sqliteDbName", "default.db", "Filename of SQLite database")
	hostingPort = flag.Int("hostingPort", 8088, "Default hosting port for the service")
	mySQLConnection = flag.String("mySQLConnection", "default", "Standard connection for mysql database")
}

func cleanupTests(){
	os.Remove(*sqliteDbName)
}

func TestStartService(t *testing.T){
	initTestingFlags()
	go StartService()

	time.Sleep(10* time.Millisecond) //hack to wait till service starts.
	url := "http://localhost:8088/health"

	//call http, check response
	response, err := http.Get(url)
	if err != nil {
		t.Error(err)
	}

	if response == nil{
		t.Error("Empty response")
	}

	//stop service, delete temp db file.
	cleanupTests()
}


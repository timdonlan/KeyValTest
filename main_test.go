package main

import (
	"flag"
	"net/http"
	"testing"

	"fmt"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"os"
	"time"
)

func initTestingFlags() {
	sqliteDbName = flag.String("sqliteDbName", "default.db", "Filename of SQLite database")
	hostingPort = flag.Int("hostingPort", 8088, "Default hosting port for the service")
	mySQLConnection = flag.String("mySQLConnection", "default", "Standard connection for mysql database")
}

func cleanupTests() {
	os.Remove(*sqliteDbName)
}

func TestFoo(t *testing.T) {
	handler := func(c *gin.Context) {
		c.String(http.StatusOK, "bar")
	}

	router := gin.New()
	router.GET("/foo", handler)

	req, _ := http.NewRequest("GET", "/foo", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	fmt.Print(resp.Body.String())

	//assert.Equal(t, resp.Body.String(), "bar")
}

func TestStartService(t *testing.T) {
	initTestingFlags()
	go StartService()

	time.Sleep(10 * time.Millisecond) //hack to wait till service starts.
	url := "http://localhost:8088/health"

	//call http, check response
	response, err := http.Get(url)
	if err != nil {
		t.Error(err)
	}

	if response == nil {
		t.Error("Empty response")
	}

	//stop service, delete temp db file.
	cleanupTests()
}

/*
func TestGetKey(t *testing.T){
	//initTestingFlags()

	//insert key
	model.Connect(*sqliteDbName)
	model.CreateKeyVal("test","foo")

	go StartService()

	time.Sleep(10* time.Millisecond) //hack to wait till service starts.
	url := "http://localhost:8088/key/test"

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

*/

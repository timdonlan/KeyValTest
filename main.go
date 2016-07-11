package main

import (
	"KeyValTest/model"
	"KeyValTest/router"
	_ "KeyValTest/router"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vharitonsky/iniflags"
)

var (
	sqliteDbName *string
	hostingPort *int
	mySQLConnection  *string
)

func main() {
	InitializeFlags()
	StartService()
}

func StartService(){
	InitializeDatabase()
	router.StartService("",*hostingPort)
}

func InitializeFlags() {

	sqliteDbName   = flag.String("sqliteDbName", "default.db", "Filename of SQLite database")
	hostingPort = flag.Int("hostingPort", 8080, "Default hosting port for the service")
	mySQLConnection = flag.String("mySQLConnection", "default", "Standard connection for mysql database")

	iniflags.Parse()
}

func InitializeDatabase() {
	model.OpenDB(*sqliteDbName)
}

func testMySQL() {
	model.Connect(*mySQLConnection)
	model.InsertDSGroup()
	dsGroup := model.SelectDSGroup()

	json, err := json.Marshal(dsGroup)
	if err == nil {
		fmt.Printf("%s", json)
	}
}

func testDB() {
	dbName := "KeyVal.db"

	model.ResetDB(dbName)
	model.CreateKeyVal("hello", "world")

	model.CreateKeyVal("world2", "hello2")

	model.UpdateKeyVal("hello", "world3")
	model.DeleteKeyVal("world2")

	dataArray, err := model.GetAll()
	if err != nil {
		return
	}

	json, err := json.Marshal(dataArray)
	fmt.Printf("%s", json)

	for _, data := range dataArray {
		if data != nil {

			fmt.Printf("%s: %s\n", data.Key, data.Value)
		}
	}
}

func testGin() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and server on 0.0.0.0:8080
}

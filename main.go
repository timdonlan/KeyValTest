package main

import ("fmt"
	_ "github.com/mattn/go-sqlite3"
	"KeyValTest/model"
	_ "KeyValTest/router"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"flag"
	"github.com/vharitonsky/iniflags"
	"KeyValTest/router"
)

var(
	dbName = flag.String("dbName","default.db","Description")
	mySQLConnection = flag.String("mySQLConnection","default","Description")
	dbUser = flag.String("dbUser","user","DB Username")
	dbPass = flag.String("dbPass","password", "DB Password")
)

func main() {

	InitializeFlags()
	InitializeDatabase()
	router.StartService()

	//testMySQL()
	//model.ResetDB(*dbName)
	//model.InitDB(*dbName)
	//model.CreateKeyVal("hello", "world")
	//model.CreateKeyVal("world2", "hello2")

}

func InitializeFlags(){
	iniflags.Parse()
}

func InitializeDatabase(){
	model.OpenDB(*dbName)
}

func testMySQL(){
	model.Connect(*mySQLConnection)
	model.InsertDSGroup()
	dsGroup := model.SelectDSGroup()

	json,err := json.Marshal(dsGroup)
	if err == nil{
		fmt.Printf("%s",json)
	}
}

func testDB() {
	dbName := "KeyVal.db"

	model.ResetDB(dbName)
	model.InitDB(dbName)
	model.CreateKeyVal("hello", "world")

	model.CreateKeyVal("world2", "hello2")

	model.UpdateKeyVal("hello", "world3")
	model.DeleteKeyVal("world2")

	dataArray, err := model.GetAllKeyVal()
	if (err != nil) {
		return
	}

	json, err := json.Marshal(dataArray)
	fmt.Printf("%s", json)

	for _, data := range dataArray {
		if (data != nil) {

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

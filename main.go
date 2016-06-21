package main

import ("fmt"
	_ "github.com/mattn/go-sqlite3"
	"KeyValTest/model"
	"KeyValTest/router"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"flag"
	"github.com/vharitonsky/iniflags"
)

var(
	dbName = flag.String("dbName","default","Description")
	dbUser = flag.String("dbUser","user","DB Username")
	dbPass = flag.String("dbPass","password", "DB Password")
)

func main() {

	InitializeFlags()

	//model.ResetDB(dbName)
	model.InitDB(*dbName)
	//model.CreateKeyVal("hello", "world")
	//model.CreateKeyVal("world2", "hello2")
	router.StartService()

}

func InitializeFlags(){
	iniflags.Parse()
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

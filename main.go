package main

import ("fmt"
	_ "github.com/mattn/go-sqlite3"
	"KeyValTest/model"
	"encoding/json"
)

func main() {

	dbName := "KeyVal.db"

	model.ResetDB(dbName)
	model.InitDB(dbName)
	model.CreateKeyVal("hello", "world")

	model.CreateKeyVal("world2", "hello2")

	model.UpdateKeyVal("hello","world3")
	model.DeleteKeyVal("world2")

	dataArray,err := model.GetAllKeyVal()
	if(err != nil){
		return
	}

	json,err := json.Marshal(dataArray)
	fmt.Printf("%s", json)

	for _, data := range dataArray{
		if(data != nil) {

		fmt.Printf("%s: %s\n",data.Key, data.Value)
		}
	}
}

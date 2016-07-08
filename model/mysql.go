package model


import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type DataStoreGroup struct{
	Id int
	GroupName string
	ParentId int
	ProviderId int
}

const insertDS = "INSERT INTO dataStoreGroup (groupName, parentId, providerId) values ('testGroup2',0,0)"
const selectDS = "select id, groupName as 'groupname' ,parentId as 'parentid', providerId as 'providerid' from dataStoreGroup"
const selectDS2 = `select id, groupName as 'groupname' ,parentId as 'parentid', providerId as 'providerid' from dataStoreGroup
where id = ?`

func Connect(dataSourceName string){

	log.Printf("%s",dataSourceName)

	db, err := sqlx.Open("mysql", "godev:gopass@/godev")
	if(err!= nil){
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

func InsertDSGroup(){
	db, err := sqlx.Open("mysql", "godev:gopass@/godev")
	if(err!= nil){
		log.Panic(err)
	}

	stmt, err := db.Prepare(insertDS)
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Print(err)
	}

}

func SelectDSGroup() *[]DataStoreGroup{
	var retVal []DataStoreGroup

	db, err := sqlx.Open("mysql", "godev:gopass@/godev")
	if(err!= nil){
		log.Panic(err)
	}

	err = db.Select(&retVal,selectDS2,2)
	if err != nil{
		log.Printf("Error in query %s",err)
	}
	return &retVal;
}
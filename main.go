package main

import (
	"KeyVal/model"
	"KeyVal/router"
	"flag"
	"github.com/vharitonsky/iniflags"
)

var (
	sqliteDbName *string
	hostingPort  *int
)

func main() {
	InitializeFlags()
	StartService()
}

func StartService() {

	router.StartService("", *hostingPort,InitializeDatabase())
}

func InitializeFlags() {

	sqliteDbName = flag.String("sqliteDbName", "default.db", "Filename of SQLite database")
	hostingPort = flag.Int("hostingPort", 8080, "Default hosting port for the service")

	iniflags.Parse()
}

func InitializeDatabase() (*model.KeyValDAL) {
	return model.OpenDB(*sqliteDbName)
}

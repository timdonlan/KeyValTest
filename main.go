package main

import (
	"KeyValTest/model"
	"KeyValTest/router"
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
	InitializeDatabase()
	router.StartService("", *hostingPort)
}

func InitializeFlags() {

	sqliteDbName = flag.String("sqliteDbName", "default.db", "Filename of SQLite database")
	hostingPort = flag.Int("hostingPort", 8080, "Default hosting port for the service")

	iniflags.Parse()
}

func InitializeDatabase() {
	model.OpenDB(*sqliteDbName)
}

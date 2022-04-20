package main

import (
	"github.com/SargntSprinkles/lion-turtle-api/config"
	"github.com/SargntSprinkles/lion-turtle-api/db"
	"github.com/SargntSprinkles/lion-turtle-api/server"
)

func main() {
	config.Initialize()

	pgdb := db.PGDB()
	defer pgdb.Disconnect()
	server.Serve()
}

package main

import (
	"os"

	"github.com/SargntSprinkles/lion-turtle-api/db"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)
	if environment := os.Getenv("LIONTURTLE_ENV"); environment == "dev" {
		logrus.SetLevel(logrus.DebugLevel)
	}

	pgdb := db.PGDB()
	defer pgdb.Disconnect()
}

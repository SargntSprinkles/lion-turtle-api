package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

const (
	DEV        = "dev"
	STAGING    = "staging"
	PRODUCTION = "production"
)

var environment string = os.Getenv("LIONTURTLE_ENV")
var port string = os.Getenv("PORT")

func Initialize() {
	logrus.SetLevel(logrus.InfoLevel)
	if Env() == DEV {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.Infof("environment set to %s", Env())
}

func Env() string {
	switch environment {
	case "":
		environment = DEV
	case DEV:
	case STAGING:
	case PRODUCTION:
	default:
		logrus.Fatal("invalid environment level")
	}
	return environment
}

func Port() string {
	if port == "" {
		port = "8080"
	}
	return port
}

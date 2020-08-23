package cmd

import (
	"database/sql"
	"errors"
	"os"

	// go package for postgres
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func prepareDatabase() (*sql.DB, error) {
	db, err := preparePostgres()
	if err != nil {
		return nil, err
	}
	logrus.WithFields(
		logrus.Fields{
			"Function": "prepareDatabase()",
		}).Info("App : Database connected successfully!")
	return db, nil
}

func getServerAddr() string {
	port, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		logrus.WithFields(
			logrus.Fields{
				"Function": "getServerAddr()",
			}).Info("App : SERVER PORT environment variable required but not set")
		os.Exit(1)
	}
	addr := ":" + port
	return addr
}

//CheckEnv : Check if the environment variables are set
func CheckEnv() error {
	_, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		return errors.New("SERVER PORT environment variable required but not set")
	}
	_, ok = os.LookupEnv("VERLOOP_DSN")
	if !ok {
		return errors.New("VERLOOP_DSN environment variable required but not set")
	}
	_, ok = os.LookupEnv("VERLOOP_DEBUG")
	if !ok {
		return errors.New("VERLOOP_DEBUG environment variable required but not set")
	}
	return nil
}

//SetLogLevel : To set the level of logging
func SetLogLevel() {
	level, ok := os.LookupEnv("VERLOOP_DEBUG")

	if !ok {
		level = "trace"
	}
	if level != "true" {
		level = "trace"
	}
	// parse string, this is built-in feature of logrus
	loglevel, err := logrus.ParseLevel(level)
	if err != nil {
		loglevel = logrus.DebugLevel
	}
	// set global log level
	logrus.SetLevel(loglevel)
}

package cmd

import (
	log "github.com/sirupsen/logrus"
)

// Begin : Beginning of the app
func Begin() {
	err := CheckEnv()
	if err != nil {
		log.WithFields(
			log.Fields{
				"function": "Begin",
			}).Info("Error : ", err.Error())
		panic(err)
	}
	SetLogLevel()
	db, err := prepareDatabase()
	if err != nil {
		log.WithFields(
			log.Fields{
				"function": "Begin()",
			}).Info("Error : Database connection failed! ", err.Error())
		panic(err)
	} else {
		app := NewApp(db)
		app.Serve(getServerAddr())
	}
}

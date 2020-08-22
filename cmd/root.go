package cmd

import (
	"log"
)

// Begin : Beginning of the app
func Begin() {
	db, err := prepareDatabase()
	if err != nil {
		log.Println("App : Database connection failed!")
		panic(err)
	} else {
		app := NewApp(db)
		app.Serve(getServerAddr())
	}
}

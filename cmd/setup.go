package cmd

import (
	"database/sql"
	"os"
	"log"

	// go package for postgres
	_ "github.com/lib/pq"
)

func prepareDatabase() (*sql.DB, error) {
	db, err := preparePostgres()
	if err != nil {
		return nil, err
	}
	log.Println("App : Database connected successfully!")
	return db, nil
}

func getServerAddr() string {
	port, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		log.Println("App : SERVER environment variable required but not set")
		os.Exit(1)
	}
	addr := ":" + port
	return addr
}

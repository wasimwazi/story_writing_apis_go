package cmd

import (
	"database/sql"
	"errors"
	"os"
)

func preparePostgres() (*sql.DB, error) {
	url, err := getURL()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getURL() (string, error) {
	psqlURL, ok := os.LookupEnv("VERLOOP_DSN")
	if !ok {
		return "", errors.New("VERLOOP_DSN environment variable required but not set")
	}
	return psqlURL, nil
}

package cmd

import (
	"database/sql"
	"net/http"
	"storyapi/router"

	"github.com/sirupsen/logrus"
)

// App : Struct to represent this app
type App struct {
	router.Router
}

// NewApp : Function to get App Struct
func NewApp(db *sql.DB) *App {
	return &App{
		Router: router.NewRouter(db),
	}
}

// Serve : Function to Run API Server
func (a *App) Serve(addr string) {
	router := a.Router.Setup()
	logrus.WithFields(
		logrus.Fields{
			"Function": "getServerAddr()",
		}).Debug("App : Server is listening...")
	http.ListenAndServe(addr, router)
}

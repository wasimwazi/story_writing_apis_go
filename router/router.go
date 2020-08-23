package router

import (
	"database/sql"
	"storyapi/story"

	"github.com/go-chi/chi"
)

// Router : Basic Router
type Router interface {
	Setup() *chi.Mux
}

// ChiRouter : Router that holds DB connection
type ChiRouter struct {
	DB *sql.DB
}

// NewRouter : Returns Basic Router
func NewRouter(db *sql.DB) Router {
	return &ChiRouter{
		DB: db,
	}
}

// Setup : chi Router
func (r *ChiRouter) Setup() *chi.Mux {
	cr := chi.NewRouter()
	storyhandler := story.NewHTTPHandler(r.DB)
	cr.Post("/add", storyhandler.AddStory)
	cr.Get("/stories", storyhandler.GetStoryList)
	cr.Get("/stories/{id}", storyhandler.GetStory)
	return cr
}

package router

import (
	"database/sql"

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
	return cr
}

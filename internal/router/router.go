package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go/internal/routes"
)

func Setup() *chi.Mux {
	r := chi.NewRouter()
	setUpRoutes(r)
	return r
}

func setUpRoutes(r chi.Router) {
	r.Group(routes.GetRoutes())
}

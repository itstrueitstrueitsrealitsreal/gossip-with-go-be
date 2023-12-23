package routes

import (
	"encoding/json"
	"github.com/CVWO/sample-go-app/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		// Users routes
		r.Route("/users", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, req *http.Request) {
				response, _ := handlers.HandleListUsers(w, req)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			})

			r.Get("/{id}", func(w http.ResponseWriter, req *http.Request) {
				response, _ := handlers.HandleGetUser(w, req)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			})
		})

		// Threads routes
		r.Route("/threads", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, req *http.Request) {
				response, _ := handlers.HandleListThreads(w, req)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			})

			r.Get("/{id}", func(w http.ResponseWriter, req *http.Request) {
				response, _ := handlers.HandleGetThread(w, req)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			})
		})

		// Posts routes
		r.Route("/posts", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, req *http.Request) {
				response, _ := handlers.HandleListPosts(w, req)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			})

			r.Get("/{id}", func(w http.ResponseWriter, req *http.Request) {
				response, _ := handlers.HandleGetPost(w, req)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			})
		})

		// Tags routes
		r.Route("/tags", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, req *http.Request) {
				response, _ := handlers.HandleListTags(w, req)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			})

			r.Get("/{id}", func(w http.ResponseWriter, req *http.Request) {
				response, _ := handlers.HandleGetTag(w, req)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			})
		})

		// Add other routes as needed
	}
}

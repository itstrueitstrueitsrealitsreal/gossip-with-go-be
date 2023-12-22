package routes

import (
	"encoding/json"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/handlers/threads"
	"github.com/CVWO/sample-go-app/internal/handlers/users"
	"github.com/go-chi/chi/v5"
)

func GetRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/users", func(w http.ResponseWriter, req *http.Request) {
		response, _ := users.HandleList(w, req)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	r.Route("/threads", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, req *http.Request) {
			// Handle listing all threads
			response, _ := threads.HandleList(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})

		r.Get("/:id", func(w http.ResponseWriter, req *http.Request) {
			// Handle viewing a specific thread
			threadID := chi.URLParam(req, "id")
			response, _ := threads.HandleView(w, req, threadID)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
	})

	return r
}

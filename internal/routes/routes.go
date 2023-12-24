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

			r.Post("/", func(w http.ResponseWriter, req *http.Request) {
				response, _ := handlers.HandleCreateUser(w, req)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			})

			r.Put("/{id}", func(w http.ResponseWriter, req *http.Request) {
				response, _ := handlers.HandleUpdateUser(w, req)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			})

			r.Delete("/{id}", func(w http.ResponseWriter, req *http.Request) {
				response, _ := handlers.HandleDeleteUser(w, req)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			})
		})

		// Threads routes
		r.Route("/threads", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, req *http.Request) {
				response, _ := handlers.HandleListThreads(w, req)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response.Payload)
			})

			r.Get("/{id}", func(w http.ResponseWriter, req *http.Request) {
				response, _ := handlers.HandleGetThread(w, req)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response.Payload)
			})
			//
			//	r.Post("/", func(w http.ResponseWriter, req *http.Request) {
			//		response, _ := handlers.HandleCreateThread(w, req)
			//		w.Header().Set("Content-Type", "application/json")
			//		json.NewEncoder(w).Encode(response)
			//	})
			//
			//	r.Put("/{id}", func(w http.ResponseWriter, req *http.Request) {
			//		response, _ := handlers.HandleUpdateThread(w, req)
			//		w.Header().Set("Content-Type", "application/json")
			//		json.NewEncoder(w).Encode(response)
			//	})
			//
			//	r.Delete("/{id}", func(w http.ResponseWriter, req *http.Request) {
			//		response, _ := handlers.HandleDeleteThread(w, req)
			//		w.Header().Set("Content-Type", "application/json")
			//		json.NewEncoder(w).Encode(response)
			//	})
		})

		// Posts routes
		r.Route("/posts", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, req *http.Request) {
				response, _ := handlers.HandleListPosts(w, req)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response.Payload)
			})

			r.Get("/{id}", func(w http.ResponseWriter, req *http.Request) {
				response, _ := handlers.HandleGetPost(w, req)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response.Payload)
			})

			//	r.Post("/", func(w http.ResponseWriter, req *http.Request) {
			//		response, _ := handlers.HandleCreatePost(w, req)
			//		w.Header().Set("Content-Type", "application/json")
			//		json.NewEncoder(w).Encode(response)
			//	})
			//
			//	r.Put("/{id}", func(w http.ResponseWriter, req *http.Request) {
			//		response, _ := handlers.HandleUpdatePost(w, req)
			//		w.Header().Set("Content-Type", "application/json")
			//		json.NewEncoder(w).Encode(response)
			//	})
			//
			//	r.Delete("/{id}", func(w http.ResponseWriter, req *http.Request) {
			//		response, _ := handlers.HandleDeletePost(w, req)
			//		w.Header().Set("Content-Type", "application/json")
			//		json.NewEncoder(w).Encode(response)
			//	})
		})

		// Tags routes
		r.Route("/tags", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, req *http.Request) {
				response, _ := handlers.HandleListTags(w, req)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response.Payload)
			})

			r.Get("/{id}", func(w http.ResponseWriter, req *http.Request) {
				response, _ := handlers.HandleGetTag(w, req)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response.Payload)
			})
		})
	}
}

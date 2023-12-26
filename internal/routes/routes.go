package routes

import (
	"encoding/json"
	"net/http"

	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go-be/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func respondWithError(w http.ResponseWriter, statusCode int, errorMessage string) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(map[string]string{"error": errorMessage})
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		// Users routes
		r.Route("/users", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, req *http.Request) {
				response, err := handlers.HandleListUsers(w, req)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, err.Error())
					return
				}

				w.Header().Set("Content-Type", "application/json")
				err = json.NewEncoder(w).Encode(response)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, "Error encoding response")
				}
			})

			r.Get("/{id}", func(w http.ResponseWriter, req *http.Request) {
				response, err := handlers.HandleGetUser(w, req)
				if err != nil {
					respondWithError(w, http.StatusNotFound, "User not found")
					return
				}

				w.Header().Set("Content-Type", "application/json")
				err = json.NewEncoder(w).Encode(response)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, "Error encoding response")
				}
			})

			r.Post("/", func(w http.ResponseWriter, req *http.Request) {
				response, err := handlers.HandleCreateUser(w, req)
				if err != nil {
					respondWithError(w, http.StatusBadRequest, err.Error())
					return
				}

				w.Header().Set("Content-Type", "application/json")
				err = json.NewEncoder(w).Encode(response)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, "Error encoding response")
				}
			})

			r.Put("/{id}", func(w http.ResponseWriter, req *http.Request) {
				response, err := handlers.HandleUpdateUser(w, req)
				if err != nil {
					respondWithError(w, http.StatusBadRequest, err.Error())
					return
				}

				w.Header().Set("Content-Type", "application/json")
				err = json.NewEncoder(w).Encode(response)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, "Error encoding response")
				}
			})

			r.Delete("/{id}", func(w http.ResponseWriter, req *http.Request) {
				response, err := handlers.HandleDeleteUser(w, req)
				if err != nil {
					respondWithError(w, http.StatusNotFound, "User not found")
					return
				}

				w.Header().Set("Content-Type", "application/json")
				err = json.NewEncoder(w).Encode(response)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, "Error encoding response")
				}
			})
		})

		// Threads routes
		r.Route("/threads", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, req *http.Request) {
				response, err := handlers.HandleListThreads(w, req)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, err.Error())
					return
				}

				w.Header().Set("Content-Type", "application/json")
				err = json.NewEncoder(w).Encode(response)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, "Error encoding response")
				}
			})

			r.Get("/{id}", func(w http.ResponseWriter, req *http.Request) {
				response, err := handlers.HandleGetThread(w, req)
				if err != nil {
					respondWithError(w, http.StatusNotFound, "Thread not found")
					return
				}

				w.Header().Set("Content-Type", "application/json")
				err = json.NewEncoder(w).Encode(response)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, "Error encoding response")
				}
			})

			r.Post("/", func(w http.ResponseWriter, req *http.Request) {
				response, err := handlers.HandleCreateThread(w, req)
				if err != nil {
					respondWithError(w, http.StatusBadRequest, err.Error())
					return
				}

				w.Header().Set("Content-Type", "application/json")
				err = json.NewEncoder(w).Encode(response)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, "Error encoding response")
				}
			})

			r.Put("/{id}", func(w http.ResponseWriter, req *http.Request) {
				response, err := handlers.HandleUpdateThread(w, req)
				if err != nil {
					respondWithError(w, http.StatusBadRequest, err.Error())
					return
				}

				w.Header().Set("Content-Type", "application/json")
				err = json.NewEncoder(w).Encode(response)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, "Error encoding response")
				}
			})

			r.Delete("/{id}", func(w http.ResponseWriter, req *http.Request) {
				response, err := handlers.HandleDeleteThread(w, req)
				if err != nil {
					respondWithError(w, http.StatusNotFound, "Thread not found")
					return
				}

				w.Header().Set("Content-Type", "application/json")
				err = json.NewEncoder(w).Encode(response)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, "Error encoding response")
				}
			})
		})

		// Posts routes
		r.Route("/posts", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, req *http.Request) {
				response, err := handlers.HandleListPosts(w, req)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, err.Error())
					return
				}

				w.Header().Set("Content-Type", "application/json")
				err = json.NewEncoder(w).Encode(response)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, "Error encoding response")
				}
			})

			r.Get("/{id}", func(w http.ResponseWriter, req *http.Request) {
				response, err := handlers.HandleGetPost(w, req)
				if err != nil {
					respondWithError(w, http.StatusNotFound, "Post not found")
					return
				}

				w.Header().Set("Content-Type", "application/json")
				err = json.NewEncoder(w).Encode(response)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, "Error encoding response")
				}
			})

			r.Post("/", func(w http.ResponseWriter, req *http.Request) {
				response, err := handlers.HandleCreatePost(w, req)
				if err != nil {
					respondWithError(w, http.StatusBadRequest, err.Error())
					return
				}

				w.Header().Set("Content-Type", "application/json")
				err = json.NewEncoder(w).Encode(response)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, "Error encoding response")
				}
			})

			r.Put("/{id}", func(w http.ResponseWriter, req *http.Request) {
				response, err := handlers.HandleUpdatePost(w, req)
				if err != nil {
					respondWithError(w, http.StatusBadRequest, err.Error())
					return
				}

				w.Header().Set("Content-Type", "application/json")
				err = json.NewEncoder(w).Encode(response)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, "Error encoding response")
				}
			})

			r.Delete("/{id}", func(w http.ResponseWriter, req *http.Request) {
				response, err := handlers.HandleDeletePost(w, req)
				if err != nil {
					respondWithError(w, http.StatusNotFound, "Post not found")
					return
				}

				w.Header().Set("Content-Type", "application/json")
				err = json.NewEncoder(w).Encode(response)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, "Error encoding response")
				}
			})
		})

		// Tags routes
		r.Route("/tags", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, req *http.Request) {
				response, err := handlers.HandleListTags(w, req)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, err.Error())
					return
				}

				w.Header().Set("Content-Type", "application/json")
				err = json.NewEncoder(w).Encode(response)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, "Error encoding response")
				}
			})

			r.Get("/{id}", func(w http.ResponseWriter, req *http.Request) {
				response, err := handlers.HandleGetTag(w, req)
				if err != nil {
					respondWithError(w, http.StatusNotFound, "Tag not found")
					return
				}

				w.Header().Set("Content-Type", "application/json")
				err = json.NewEncoder(w).Encode(response)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, "Error encoding response")
				}
			})
		})
	}
}

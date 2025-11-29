package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(h *handler) http.Handler {
	r := chi.NewRouter()

	r.Route("/notes", func(r chi.Router) {
		r.Get("/", h.GetAll)
		r.Post("/", h.Create)
		r.Get("/{id}", h.Get)
		r.Put("/{id}", h.Update)
		r.Delete("/{id}", h.Delete)
	})

	return r
}

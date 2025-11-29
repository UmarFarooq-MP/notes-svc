package http

import (
	"encoding/json"
	"net/http"
	"web3/internal/domain"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func NewHandler(svc domain.Notes) *handler {
	return &handler{svc: svc}
}

type handler struct {
	svc domain.Notes
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {

	var req CreateNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if req.Title == "" || req.Content == "" {
		http.Error(w, "title and content are required", http.StatusBadRequest)
		return
	}

	err := h.svc.Create(domain.Note{
		Id:      uuid.New(),
		Content: req.Content,
		Title:   req.Title,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(domain.Note{
		Id:      uuid.New(),
		Content: req.Content,
		Title:   req.Title,
	})
}

func (h handler) Update(w http.ResponseWriter, r *http.Request) {

	var req UpdateNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if req.Title == "" || req.Content == "" {
		http.Error(w, "title and content are required", http.StatusBadRequest)
		return
	}

	err := h.svc.Update(req.ID, domain.Note{
		Id:      uuid.New(),
		Content: req.Content,
		Title:   req.Title,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(domain.Note{
		Id:      uuid.New(),
		Content: req.Content,
		Title:   req.Title,
	})
}

func (h handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if uuidParse, err := uuid.Parse(id); err == nil {
		err = h.svc.Delete(uuidParse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return
	}
	http.Error(w, "invalid uuid", http.StatusBadRequest)
}

func (h handler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid uuid", http.StatusBadRequest)
		return
	}

	note, err := h.svc.Get(id)
	if err != nil {
		if err.Error() == "note not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(note)
}

func (h handler) GetAll(w http.ResponseWriter, r *http.Request) {
	notes, err := h.svc.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}

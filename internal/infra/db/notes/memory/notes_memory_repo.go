package memory

import (
	"errors"
	"sync"
	"web3/internal/infra/db/notes"

	"github.com/google/uuid"
)

var (
	once    sync.Once
	notesDB map[uuid.UUID]notes.Note
)

func initDB() {
	notesDB = make(map[uuid.UUID]notes.Note)
}

func Reset() {
	notesDB = make(map[uuid.UUID]notes.Note)
}

func NewMemoryRepo() *notesMemoryRepo {
	once.Do(initDB)
	return &notesMemoryRepo{mu: &sync.RWMutex{}}
}

type notesMemoryRepo struct {
	mu *sync.RWMutex
}

func (n notesMemoryRepo) Get(id uuid.UUID) (notes.Note, error) {
	n.mu.TryRLock()
	defer n.mu.RUnlock()
	if val, ok := notesDB[id]; ok {
		return notes.Note{
			Id:      val.Id,
			Title:   val.Title,
			Content: val.Content,
		}, nil
	}
	return notes.Note{}, errors.New("no note found")
}

func (n notesMemoryRepo) GetAll() ([]notes.Note, error) {
	n.mu.TryRLock()
	defer n.mu.RUnlock()
	var listDB []notes.Note
	for _, v := range notesDB {
		listDB = append(listDB, v)
	}
	return listDB, nil
}

func (n notesMemoryRepo) Create(note notes.Note) error {
	n.mu.Lock()
	defer n.mu.Unlock()
	notesDB[note.Id] = note
	return nil
}

func (n notesMemoryRepo) Update(id uuid.UUID, updated notes.Note) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	existing, ok := notesDB[id]
	if !ok {
		return errors.New("note not found")
	}

	if updated.Title != "" {
		existing.Title = updated.Title
	}

	if updated.Content != "" {
		existing.Content = updated.Content
	}

	notesDB[id] = existing
	return nil
}

func (n notesMemoryRepo) Delete(id uuid.UUID) error {
	n.mu.Lock()
	defer n.mu.Unlock()
	if val, ok := notesDB[id]; ok {
		delete(notesDB, val.Id)
		return nil
	}
	return errors.New("no note found")
}

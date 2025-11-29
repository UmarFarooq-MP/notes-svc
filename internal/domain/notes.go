package domain

import "github.com/google/uuid"

type Note struct {
	Id      uuid.UUID
	Title   string
	Content string
}

type Notes interface {
	Get(id uuid.UUID) (Note, error)
	GetAll() ([]Note, error)
	Create(note Note) error
	Update(id uuid.UUID, updated Note) error
	Delete(id uuid.UUID) error
}

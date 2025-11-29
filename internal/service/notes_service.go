package service

import (
	"web3/internal/domain"
	"web3/internal/infra/db/notes"

	"github.com/google/uuid"
)

func New(repo notes.Repository) *Notes {
	return &Notes{repo: repo}
}

type Notes struct {
	repo notes.Repository
}

func (n Notes) Get(id uuid.UUID) (domain.Note, error) {
	note, err := n.repo.Get(id)
	if err != nil {
		return domain.Note{}, err
	}
	return dbToDomain(note), nil
}

func (n Notes) GetAll() ([]domain.Note, error) {
	notes, err := n.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return dbToDomainAll(notes), nil
}

func (n Notes) Create(note domain.Note) error {
	return n.repo.Create(DomainToDB(note))
}

func (n Notes) Update(id uuid.UUID, updated domain.Note) error {
	return n.repo.Update(id, DomainToDB(updated))
}

func (n Notes) Delete(id uuid.UUID) error {
	return n.repo.Delete(id)
}

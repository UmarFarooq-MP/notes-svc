package service

import (
	"web3/internal/domain"
	db "web3/internal/infra/db/notes"
)

func dbToDomain(dbNote db.Note) domain.Note {
	return domain.Note{
		Id:      dbNote.Id,
		Title:   dbNote.Title,
		Content: dbNote.Content,
	}
}

func dbToDomainAll(Notes []db.Note) []domain.Note {
	var allNotes []domain.Note
	for _, dbNote := range Notes {
		allNotes = append(allNotes, domain.Note{
			Id:      dbNote.Id,
			Title:   dbNote.Title,
			Content: dbNote.Content,
		})
	}
	return allNotes
}

func DomainToDB(domainNote domain.Note) db.Note {
	return db.Note{
		Id:      domainNote.Id,
		Title:   domainNote.Title,
		Content: domainNote.Content,
	}
}

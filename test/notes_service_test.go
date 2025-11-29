package test

import (
	"testing"

	"web3/internal/domain"
	memoryRepo "web3/internal/infra/db/notes/memory"
	"web3/internal/service"

	"github.com/google/uuid"
)

func TestCreateNote(t *testing.T) {
	memoryRepo.Reset()
	repo := memoryRepo.NewMemoryRepo()
	svc := service.New(repo)

	note := domain.Note{
		Id:      uuid.New(),
		Title:   "Hello",
		Content: "World",
	}

	err := svc.Create(note)
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	got, err := svc.Get(note.Id)
	if err != nil {
		t.Fatalf("failed fetching after create: %v", err)
	}

	if got.Title != note.Title {
		t.Errorf("expected title %s, got %s", note.Title, got.Title)
	}
}

func TestGetAll(t *testing.T) {
	memoryRepo.Reset()
	repo := memoryRepo.NewMemoryRepo()
	svc := service.New(repo)

	svc.Create(domain.Note{Id: uuid.New(), Title: "A", Content: "1"})
	svc.Create(domain.Note{Id: uuid.New(), Title: "B", Content: "2"})

	notes, err := svc.GetAll()
	if err != nil {
		t.Fatalf("get all failed: %v", err)
	}

	if len(notes) != 2 {
		t.Fatalf("expected 2 notes, got %d", len(notes))
	}
}

func TestGetNoteNotFound(t *testing.T) {
	repo := memoryRepo.NewMemoryRepo()
	svc := service.New(repo)

	fakeID := uuid.New()

	_, err := svc.Get(fakeID)
	if err == nil {
		t.Fatalf("expected error for missing id")
	}
}

func TestUpdateNote(t *testing.T) {
	memoryRepo.Reset()
	repo := memoryRepo.NewMemoryRepo()
	svc := service.New(repo)

	id := uuid.New()

	svc.Create(domain.Note{
		Id:      id,
		Title:   "Old",
		Content: "Old Content",
	})

	updated := domain.Note{
		Title:   "New",
		Content: "New Content",
	}

	svc.Update(id, updated)

	got, err := svc.Get(id)
	if err != nil {
		t.Fatalf("fetch after update failed: %v", err)
	}

	if got.Title != "New" {
		t.Fatalf("expected New, got %s", got.Title)
	}
}

func TestDeleteNote(t *testing.T) {
	repo := memoryRepo.NewMemoryRepo()
	svc := service.New(repo)

	id := uuid.New()

	svc.Create(domain.Note{
		Id:      id,
		Title:   "Test",
		Content: "Delete Me",
	})

	err := svc.Delete(id)
	if err != nil {
		t.Fatalf("delete failed: %v", err)
	}

	_, err = svc.Get(id)
	if err == nil {
		t.Fatalf("expected error after delete")
	}
}

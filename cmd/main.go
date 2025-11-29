package main

import (
	"fmt"
	goHttp "net/http"
	memoryRepo "web3/internal/infra/db/notes/memory"
	"web3/internal/infra/http"
	"web3/internal/service"
)

func main() {
	notesSvc := service.New(memoryRepo.NewMemoryRepo())
	notesHandler := http.NewHandler(notesSvc)
	router := http.NewRouter(notesHandler)

	fmt.Println("Server started on http://localhost:8080")

	if err := goHttp.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error running server:", err)
	}
}

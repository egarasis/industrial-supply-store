package main

import (
	"log"

	"industrial-supply-store/internal/config"
	"industrial-supply-store/internal/handlers"
	dbrepo "industrial-supply-store/internal/repository/db"
	"industrial-supply-store/internal/usecase"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := dbrepo.NewUserRepository(db)

	uc := usecase.NewUserUsecase(repo)

	handler := handlers.NewUserHandler(uc)

	handler.Run()
}
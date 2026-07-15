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

	// Repository
	repoUser := dbrepo.NewUserRepository(db)
	repoProduct := dbrepo.NewProductRepository(db)

	// Usecase
	ucUser := usecase.NewUserUsecase(repoUser)
	ucAdmin := usecase.NewAdminUsecase(repoProduct)

	// handler
	handlerAdmin := handlers.NewAdminHandler(ucAdmin)
	handler := handlers.NewUserHandler(ucUser, handlerAdmin)

	handler.Run()
}

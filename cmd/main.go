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
	repoOrder := dbrepo.NewOrderRepository(db)

	// Usecase
	ucUser := usecase.NewUserUsecase(repoUser)
<<<<<<< HEAD
	ucAdmin := usecase.NewAdminUsecase(repoProduct)
	ucCustomer := usecase.NewCustomerUsecase(db, repoOrder, repoProduct, repoUser)
=======
	ucAdmin := usecase.NewAdminUsecase(repoProduct, repoOrder)
	ucCustomer := usecase.NewCustomerUsecase(db, repoOrder, repoProduct)
>>>>>>> a9d2308fdc4245458fd69dd2b7b286b0217a42fc

	// handler
	handlerAdmin := handlers.NewAdminHandler(ucAdmin, ucUser)
	handlerCustomer := handlers.NewCustomerHandler(ucCustomer)
	handler := handlers.NewUserHandler(ucUser, handlerAdmin, handlerCustomer)

	handler.Run()
}

package usecase

import (
	"industrial-supply-store/internal/domain"
)

type adminUsecase struct {
	productRepo domain.ProductRepository
}

func NewAdminUsecase(productRepo domain.ProductRepository) domain.AdminUsecase {
	return &adminUsecase{
		productRepo: productRepo,
	}
}
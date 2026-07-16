package usecase

import (
	"context"
	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: repo,
	}
}

// Implementasi semua fungsi agar sesuai dengan interface UserUsecase

func (u *userUsecase) UpdateProfile(ctx context.Context, profile entity.UserProfile) error {
	return u.userRepo.UpdateProfile(ctx, profile)
}

func (u *userUsecase) Login(email, password string) (*entity.User, error) {
	// Panggil repository untuk cek user
	return u.userRepo.FindByEmail(context.Background(), email)
}

func (u *userUsecase) Register(email, password, role string) error {
	user := &entity.User{
		Email:    email,
		Password: password,
		Role:     role,
	}
	_, err := u.userRepo.Create(context.Background(), user)
	return err
}

func (u *userUsecase) GetAll() ([]entity.User, error) {
	return u.userRepo.FindAll(context.Background())
}

func (u *userUsecase) GetByID(id int) (*entity.User, error) {
	return u.userRepo.FindByID(context.Background(), id)
}

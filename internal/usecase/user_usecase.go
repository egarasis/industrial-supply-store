package usecase

import (
	"context"
	"fmt"
	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (u *userUsecase) Register(email, password, role string) error {
	ctx := context.Background()

	// Validate role
	role = strings.ToLower(strings.TrimSpace(role))
	if role != entity.RoleAdmin &&
		role != entity.RoleCustomer {

		fmt.Println("Invalid role")
		return fmt.Errorf("Invalid role")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Failed to hash password")
		return err
	}

	user := &entity.User{
		Email:    email,
		Role:     role,
		Password: string(hash),
	}

	_, err = u.repo.Create(ctx, user)

	return err
}

func (u *userUsecase) Login(email, password string) (*entity.User, error) {
	ctx := context.Background()

	user, err := u.repo.FindByEmail(ctx, email)

	if err != nil {
		fmt.Println("Email not found")
		return nil, fmt.Errorf("Email not found")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		fmt.Println("Wrong Password")
		return nil, fmt.Errorf("Wrong Password")
	}

	return user, nil
}

func (u *userUsecase) GetAll() ([]entity.User, error) {
	ctx := context.Background()
	return u.repo.FindAll(ctx)
}

func (u *userUsecase) GetByID(id int) (*entity.User, error) {
	ctx := context.Background()

	return u.repo.FindByID(ctx, id)
}
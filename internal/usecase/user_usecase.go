package usecase

import (
	"context"
	"fmt" // Ditambahkan untuk fmt.Println & fmt.Errorf
	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: repo,
	}
}

// 1. Fungsi Register dengan validasi role dan enkripsi password bcrypt
func (u *userUsecase) Register(email, password, role string) error {
	// Validate role
	role = strings.ToLower(strings.TrimSpace(role))
	if role != entity.RoleAdmin && role != entity.RoleCustomer {
		fmt.Println("Invalid role")
		return fmt.Errorf("invalid role")
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

	_, err = u.userRepo.Create(context.Background(), user)
	return err
}

func (u *userUsecase) Login(email, password string) (*entity.User, error) {
	return u.userRepo.FindByEmail(context.Background(), email)
}

func (u *userUsecase) GetAll() ([]entity.User, error) {
	return u.userRepo.FindAll(context.Background())
}

func (u *userUsecase) GetByID(id int) (*entity.User, error) {
	return u.userRepo.FindByID(context.Background(), id)
}

// 2. Tambahkan implementasi fungsi UpdateProfile agar bisa dipanggil oleh handler
func (u *userUsecase) UpdateProfile(ctx context.Context, profile entity.UserProfile) error {
	return u.userRepo.UpdateProfile(ctx, profile)
}

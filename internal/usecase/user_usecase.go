package usecase

import (
	"context"
	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
<<<<<<< HEAD
=======
	"strings"

	"golang.org/x/crypto/bcrypt"
>>>>>>> a9d2308fdc4245458fd69dd2b7b286b0217a42fc
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

<<<<<<< HEAD
func (u *userUsecase) UpdateProfile(ctx context.Context, profile entity.UserProfile) error {
	return u.userRepo.UpdateProfile(ctx, profile)
=======
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
>>>>>>> a9d2308fdc4245458fd69dd2b7b286b0217a42fc
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

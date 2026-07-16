package domain

import (
	"context"
	"industrial-supply-store/internal/model/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	FindAll(ctx context.Context) ([]entity.User, error)
	FindByID(ctx context.Context, id int) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int) error
}

type UserUsecase interface {
	Login(email, password string) (*entity.User, error)
	Register(email, password, role string) error
	GetAll() ([]entity.User, error)
	GetByID(id int) (*entity.User, error)
}

type AuthUsecase interface {
	Register(email, password, role string) error
	Login() ([]entity.User, error)
	GetByID(id int) (*entity.User, error)
	Update(id int, name, email string) error
	Delete(id int) error
}

type UserHandler interface {
	Run()
}

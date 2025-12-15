
package Usecases

import (
	"fmt"

	"github.com/Ghaby-X/task_manager/Domain"
	"github.com/Ghaby-X/task_manager/Infrastructure"
	"github.com/Ghaby-X/task_manager/Repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase interface {
	CreateUser(user Domain.User) error
	GetUserByUsername(username string) (Domain.User, error)
	PromoteUser(id primitive.ObjectID) error
	Login(username, password string) (Domain.User, error)
}

type userUsecase struct {
	userRepo        Repositories.UserRepository
	passwordService Infrastructure.PasswordService
}

func NewUserUsecase(userRepo Repositories.UserRepository, passwordService Infrastructure.PasswordService) UserUsecase {
	return &userUsecase{userRepo: userRepo, passwordService: passwordService}
}

func (u *userUsecase) CreateUser(user Domain.User) error {
	isDbEmpty, err := u.userRepo.IsDatabaseEmpty()
	if err != nil {
		return err
	}

	if isDbEmpty {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	hashedPassword, err := u.passwordService.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = hashedPassword
	
	return u.userRepo.CreateUser(user)
}

func (u *userUsecase) GetUserByUsername(username string) (Domain.User, error) {
	return u.userRepo.GetUserByUsername(username)
}

func (u *userUsecase) PromoteUser(id primitive.ObjectID) error {
	return u.userRepo.PromoteUser(id)
}

func (u *userUsecase) Login(username, password string) (Domain.User, error) {
	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return Domain.User{}, err
	}

	if !u.passwordService.CheckPasswordHash(password, user.Password) {
		return Domain.User{}, fmt.Errorf("invalid credentials")
	}
	return user, nil
}

package services

import (
	"fmt"
	"log"

	"github.com/dfriveraa/glowing-octo-memory/app/internal/adapters/repositories"
	"github.com/dfriveraa/glowing-octo-memory/app/internal/core/domain"
)

type UserService struct {
	repo            repositories.UserRepo
	securityService passwordService
}

func NewUserService(dbClient repositories.Db) *UserService {
	return &UserService{
		repo:            *repositories.NewUserRepo(dbClient),
		securityService: *newPasswordService(),
	}
}

func (us *UserService) CreateUser(newUser *domain.User) error {
	var err error
	newUser.Password, err = us.securityService.hashPassword(newUser.Password)
	if err != nil {
		log.Println(err)
		return err
	}
	err = us.repo.Create(newUser)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (us *UserService) GetUserById(userId int) (*domain.User, error) {
	user, err := us.repo.GetById(userId)
	if err != nil {
		log.Println(err)
	}
	return user, err
}

func (us *UserService) Authenticate(email string, plainPassword string) (*domain.User, error) {
	user, err := us.repo.GetByEmail(email)
	if err != nil {
		log.Println(err)
		return nil, err

	}
	successAuth := us.securityService.comparePassword(plainPassword, user.Password)
	if successAuth {
		return user, nil
	} else {
		return nil, fmt.Errorf("invalid credentials")
	}
}


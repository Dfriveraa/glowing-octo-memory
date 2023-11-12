package services

import (
	"fmt"
	"log"
	"mime/multipart"

	"github.com/dfriveraa/glowing-octo-memory/app/internal/adapters/infra"
	"github.com/dfriveraa/glowing-octo-memory/app/internal/adapters/repositories"
	config "github.com/dfriveraa/glowing-octo-memory/app/internal/core"
	"github.com/dfriveraa/glowing-octo-memory/app/internal/core/domain"
)

type UserService struct {
	repo            repositories.UserRepo
	securityService PasswordService
	bucketService   infra.BucketClient
}

func NewUserService(dbClient repositories.Db) *UserService {
	return &UserService{
		repo:            *repositories.NewUserRepo(dbClient),
		securityService: *NewPasswordService(),
		bucketService:   *infra.NewBucketBasics(),
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

func (us *UserService) Authenticate(email string, plainPassword string) (string, error) {
	user, err := us.repo.GetByEmail(email)
	if err != nil {
		log.Println(err)
		return "", err

	}
	successAuth := us.securityService.comparePassword(plainPassword, user.Password)
	if successAuth {
		token := us.securityService.generateToken(user)
		return token, nil
	} else {
		return "", fmt.Errorf("Invalid credentials")
	}
}

func (us *UserService) UploadProfilePicture(userId int, extension string, file *multipart.FileHeader) error {

	err := us.bucketService.UploadFile(config.Settings.BucketName, fmt.Sprintf("profiles/%v.%v", userId, extension), file)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (us *UserService) GetProfilePicture(userId int) (string, error) {
	url, err := us.bucketService.GetObjectURL(config.Settings.BucketName, fmt.Sprintf("profiles/%v.png", userId), 10000)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return url, nil
}

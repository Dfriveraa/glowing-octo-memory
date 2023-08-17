package repositories

import (
	"fmt"

	"github.com/dfriveraa/glowing-octo-memory/app/internal/core/domain"
)

type UserRepo struct {
	dbInstance Db
}

func NewUserRepo(client Db) *UserRepo {
	return &UserRepo{
		dbInstance: client,
	}
}
func (ur *UserRepo) Create(u *domain.User) error {
	result := ur.dbInstance.client.Create(&u)
	if result.RowsAffected != 1 {
		err := fmt.Errorf("Exception launched %v", result.Error)
		fmt.Print(err)
		return err
	}
	return nil
}

func (ur *UserRepo) GetById(uId int) (*domain.User, error) {
	user := domain.User{}
	result := ur.dbInstance.client.First(&user, uId)
	if result.RowsAffected != 1 {
		err := fmt.Errorf("Exception launched %v", result.Error)
		fmt.Print(err)
		return &user, err
	}
	return &user, nil
}

func (ur *UserRepo) GetByEmail(email string) (*domain.User, error) {
	user := domain.User{}
	result := ur.dbInstance.client.Where("email = ?", email).First(&user)
	if result.RowsAffected != 1 {
		err := fmt.Errorf("Exception launched %v", result.Error)
		fmt.Print(err)
		return &user, err
	}
	return &user, nil
}

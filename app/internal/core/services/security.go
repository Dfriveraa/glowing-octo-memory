package services

import (
	"golang.org/x/crypto/bcrypt"
)

type passwordService struct {
}

func newPasswordService() *passwordService {
	return &passwordService{}
}

func (s *passwordService) hashPassword(password string) (string, error) {
	// generate a new salt
	salt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	// concatenate the salt and hashed password
	return string(salt), nil
}

func (s *passwordService) comparePassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

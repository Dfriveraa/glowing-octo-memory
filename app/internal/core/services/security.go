package services

import (
	"fmt"
	"time"

	config "github.com/dfriveraa/glowing-octo-memory/app/internal/core"
	"github.com/dfriveraa/glowing-octo-memory/app/internal/core/domain"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type canDispatch struct {
	CanDispatchLTL bool `json:"can_dispatch_ltl"`
	CanDispatchFTL bool `json:"can_dispatch_ftl"`
	CanDispatchEXP bool `json:"can_dispatch_exp"`
}

type CustomClaims struct {
	Name         string          `json:"name"`
	Role         string          `json:"role"`
	CanDispatch  canDispatch     `json:"can_dispatch"`
	CanDispatch2 map[string]bool `json:"can_dispatch2"`
	jwt.RegisteredClaims
}
type PasswordService struct {
}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (s *PasswordService) hashPassword(password string) (string, error) {
	// generate a new salt
	salt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	// return the hashed password as a string
	return string(salt), nil
}

func (s *PasswordService) comparePassword(password string, hashedPassword string) bool {
	// compare the password with the hash and see if they match
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (s *PasswordService) generateToken(user *domain.User) string {

	myClaim := CustomClaims{
		Name:         user.Name,
		Role:         user.Role,
		CanDispatch:  s.fakeCanDispatch2(user),
		CanDispatch2: s.fakeCanDispatch(user),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "glowing-octo-memory",
			Subject:   user.Email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)
	tokenPlain, _ := token.SignedString([]byte(config.Settings.JWTSecret))
	return tokenPlain
}

func (s *PasswordService) fakeCanDispatch(user *domain.User) map[string]bool {
	if user.Role == "admin" {
		return map[string]bool{
			"can_dispatch_ltl": true,
			"can_dipatch_ftl":  true,
			"can_dispatch_exp": true,
		}
	} else {
		return map[string]bool{
			"can_dispatch_ltl": false,
			"can_dipatch_ftl":  false,
			"can_dispatch_exp": false,
		}
	}
}

func (s *PasswordService) fakeCanDispatch2(user *domain.User) canDispatch {
	if user.Role == "admin" {
		return canDispatch{
			CanDispatchLTL: true,
			CanDispatchFTL: true,
			CanDispatchEXP: true,
		}
	} else {
		return canDispatch{
			CanDispatchLTL: false,
			CanDispatchFTL: false,
			CanDispatchEXP: false,
		}
	}
}

func (s *PasswordService) validateSignedString(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	// return the secret key
	return []byte(config.Settings.JWTSecret), nil
}
func (s *PasswordService) GetCurrentUser(t string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(t, &CustomClaims{}, s.validateSignedString)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}

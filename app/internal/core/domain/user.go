package domain

import "time"

type User struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Role      string    `json:"role"`
	Email     string    `json:"email" gorm:"uniqueIndex"`
	Password  string    `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Role    string `json:"role"`
	Email   string `json:"email"`
}

func HidePassword(user *User) *UserResponse {
	return &UserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Surname: user.Surname,
		Role:    user.Role,
		Email:   user.Email,
	}
}

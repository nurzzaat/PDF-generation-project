package models

import "context"

type User struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	CreatedAt string `json:"createdAt"`
}

type UserRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Password struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UserRepository interface {
	GetUserByEmail(c context.Context, email string) (User, error)
	GetUserByID(c context.Context, userID int) (User, error)
	GetProfile(c context.Context, userID int) (User, error)

	CreateUser(c context.Context, user UserRequest) (int, error)
	EditUser(c context.Context, user User) (int, error)
	DeleteUser(c context.Context, userID int) error
	SetUserPassword(c context.Context, password string, userID int) error
}

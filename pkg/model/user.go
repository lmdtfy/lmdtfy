package model

import (
	"errors"
	"strings"

	"code.google.com/p/go.crypto/bcrypt"
)

var (
	ErrInvalidPassword = errors.New("Username or password incorrect.")
)

type User struct {
	ID             string
	Email          string
	EmailLower     string
	HashedPassword string
	IsAdmin        bool
}

func NewUser(email, password, role string) (*User, error) {

	emailLower := strings.ToLower(email)

	user := User{
		Email:      email,
		EmailLower: emailLower,
		IsAdmin:    false,
	}

	err := user.SetPassword(password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) SetPassword(password string) error {
	// Password validation.
	switch {
	case len(password) < 6:
		return ErrInvalidPassword
	case len(password) > 265:
		return ErrInvalidPassword
	}
	// Hash password
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.HashedPassword = string(b)

	return nil
}
func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
}

package database

import (
	"errors"

	r "github.com/dancannon/gorethink"

	m "github.com/lmdtfy/lmdtfy/pkg/model"
)

var (
	//ErrUserOrPasswdIncorrect err for password incorrect
	ErrUserOrPasswdIncorrect = errors.New("Email or password incorrect.")

	//ErrUserAlreadyExists err for duplicate user
	ErrUserAlreadyExists = errors.New("User with email already exists.")

	//ErrUserPasswordRequired err for trying to save without a password.
	// TODO: Remove this after refactoring.
	ErrUserPasswordRequired = errors.New("Password is required")
)

// UserStore used to interact with db users
type UserStore struct {
	DefaultStore
}

// NewUserStore returns a UserStore.
func NewUserStore() UserStore {
	return UserStore{DefaultStore: NewDefaultStore("users")}
}

func userExist(email string) bool {
	row, _ := r.Table("users").Filter(r.Row.Field("email").Eq(email)).Run(sess)

	return !row.IsNil()
}

// Store saves a user into the db
func (us *UserStore) Store(u *m.User) error {
	if userExist(u.Email) {
		return ErrUserAlreadyExists
	}
	res, err := r.Table("users").Insert(u).RunWrite(sess)
	if err != nil {
		return err
	}
	if u.ID == "" && len(res.GeneratedKeys) == 1 {
		u.ID = res.GeneratedKeys[0]
	}

	return nil
}

// FindByEmail finds a single user that matches an email.
func (us *UserStore) FindByEmail(email string) (m.User, error) {
	var u m.User

	res, err := r.Table("users").Filter(r.Row.Field("email").Eq(email)).Run(sess)
	if err != nil {
		return u, err
	}

	err = res.One(&u)
	return u, nil

}

// GetUserByID get a user by a ID.
func GetUserByID(id string) (m.User, error) {
	u := m.User{}

	res, err := r.Table("users").Get(id).Run(sess)
	if err != nil {
		return u, err
	}

	err = res.All(&u)
	if err != nil {
		return u, err
	}

	return u, nil
}

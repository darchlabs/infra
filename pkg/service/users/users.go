package userservice

import (
	"errors"

	userstorage "github.com/darchlabs/infra/pkg/storage/users"
	"github.com/darchlabs/infra/pkg/users"
	"golang.org/x/crypto/bcrypt"
)

// UserService ...
type UserService struct {
	UserStore *userstorage.UserStore
}

// NewService ...
func NewService(store *userstorage.UserStore) *UserService {
	return &UserService{
		UserStore: store,
	}
}

// GetByID ...
func (us *UserService) GetByID(id string) (*users.User, error) {
	// validate id param
	if id == "" {
		return nil, errors.New("invalid id")
	}

	// get user from storage
	return us.UserStore.Get(&users.Query{
		ID: id,
	})
}

// GetByEmail ...
func (us *UserService) GetByEmail(email string) (*users.User, error) {
	// validate email param
	if email == "" {
		return nil, errors.New("invalid email")
	}

	// get user from storage
	return us.UserStore.Get(&users.Query{
		Email: email,
	})
}

// Create ...
func (us *UserService) Create(u *users.User) error {
	// validate user existence
	if u == nil {
		return errors.New("invalid user")
	}

	// validate user params
	if u.Name == "" || u.Email == "" || u.Password == "" {
		return errors.New("invalid user params")
	}

	// create user in storage
	return us.UserStore.Create(u)
}

// Update ...
func (us *UserService) Update(u *users.User) error {
	// validate user existence
	if u == nil {
		return errors.New("invalid user")
	}

	// update user from storage
	return us.UserStore.Update(u)
}

// Delete ...
func (us *UserService) Delete(u *users.User) error {
	// validate user existence
	if u == nil {
		return errors.New("invalid user")
	}

	// delete user from storage
	return us.UserStore.Delete(u)
}

// List ...
func (us *UserService) List() ([]*users.User, error) {
	// list users from storage
	return us.UserStore.List()
}

// VerifyPassword ...
func (us *UserService) VerifyPassword(email string, password string) error {
	// validate email param
	if email == "" {
		return errors.New("invalid email")
	}

	// validate password param
	if password == "" {
		return errors.New("invalid password")
	}

	// get user by email
	user, err := us.GetByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	// compare current password with hash in storage
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

package userstorage

import (
	"github.com/darchlabs/infra/pkg/storage"
)

type UserStore struct {
	Store *storage.Store
}

func New(s *storage.Store) *UserStore {
	return &UserStore{
		Store: s,
	}
}

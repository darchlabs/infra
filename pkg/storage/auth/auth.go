package authstorage

import (
	"github.com/darchlabs/infra/pkg/storage"
)

type AuthStore struct {
	Store *storage.Store
}

func New(s *storage.Store) *AuthStore {
	return &AuthStore{
		Store: s,
	}
}

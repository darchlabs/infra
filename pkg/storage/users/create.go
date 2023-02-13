package userstorage

import (
	"github.com/Masterminds/squirrel"
	"github.com/darchlabs/infra/pkg/users"
)

// Create ...
func (us *UserStore) Create(u *users.User) error {
	sql, args, err := squirrel.
		Insert("users").
		Columns("email", "name", "password").
		Values(u.Email, u.Name, u.Password).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	row := us.Store.DB.QueryRowx(sql, args...)
	if err := row.StructScan(u); err != nil {
		return err
	}

	return nil
}

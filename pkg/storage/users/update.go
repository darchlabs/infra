package userstorage

import (
	"github.com/Masterminds/squirrel"
	"github.com/darchlabs/infra/pkg/users"
)

// Update ...
func (us *UserStore) Update(u *users.User) error {
	sql, args, err := squirrel.Update("users").Set("email", u.Email).Set("name", u.Name).Set("password", u.Password).Set("verified", u.Verified).Where("id = ?", u.ID).Suffix("returning *").PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return err
	}

	row := us.Store.DB.QueryRowx(sql, args...)
	return row.StructScan(u)
}

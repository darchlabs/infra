package userstorage

import (
	"github.com/Masterminds/squirrel"
	"github.com/darchlabs/infra/pkg/users"
)

// List ...
func (us *UserStore) List() ([]*users.User, error) {
	query := squirrel.Select("*").From("users").Where("deleted_at is null")

	sql, args, err := query.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := us.Store.DB.Queryx(sql, args...)
	if err != nil {
		return nil, err
	}

	uu := make([]*users.User, 0)

	for rows.Next() {
		u := &users.User{}
		if err := rows.StructScan(u); err != nil {
			return nil, err
		}

		uu = append(uu, u)
	}

	return uu, nil
}

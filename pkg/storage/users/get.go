package userstorage

import (
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/darchlabs/infra/pkg/users"
)

// Get ...
func (us *UserStore) Get(q *users.Query) (*users.User, error) {
	query := squirrel.Select("*").From("users").Where("deleted_at is null")

	if q.ID == "" && q.Email == "" {
		return nil, errors.New("must proovide a query")
	}

	if q.ID != "" {
		query = query.Where("id = ?", q.ID)
	}

	if q.Email != "" {
		query = query.Where("email = ?", q.Email)
	}

	sql, args, err := query.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	row := us.Store.DB.QueryRowx(sql, args...)

	c := &users.User{}
	if err := row.StructScan(c); err != nil {
		return nil, err
	}

	return c, nil
}

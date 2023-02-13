package authstorage

import (
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/darchlabs/infra/pkg/auth"
)

// Get ...
func (as *AuthStore) Get(q *auth.Query) (*auth.Auth, error) {
	query := squirrel.Select("*").From("auth").Where("deleted_at is null")

	if q.Email == "" && q.Token == "" && q.UserID == "" {
		return nil, errors.New("must proovide a query")
	}

	if q.Email != "" {
		query = query.Where("email = ?", q.Token)
	}

	if q.Token != "" {
		query = query.Where("token = ?", q.Token)
	}

	if q.UserID != "" {
		query = query.Where("user_id = ?", q.UserID)
	}

	sql, args, err := query.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	row := as.Store.DB.QueryRowx(sql, args...)

	c := &auth.Auth{}
	if err := row.StructScan(c); err != nil {
		return nil, err
	}

	return c, nil
}

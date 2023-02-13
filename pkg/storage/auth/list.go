package authstorage

import (
	"github.com/Masterminds/squirrel"
	"github.com/darchlabs/infra/pkg/auth"
)

// List ...
func (as *AuthStore) List() ([]*auth.Auth, error) {
	query := squirrel.Select("*").From("auth").Where("deleted_at is null")

	sql, args, err := query.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := as.Store.DB.Queryx(sql, args...)
	if err != nil {
		return nil, err
	}

	aa := make([]*auth.Auth, 0)

	for rows.Next() {
		a := &auth.Auth{}
		if err := rows.StructScan(a); err != nil {
			return nil, err
		}

		aa = append(aa, a)
	}

	return aa, nil
}

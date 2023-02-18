package authstorage

import (
	"github.com/Masterminds/squirrel"
	"github.com/darchlabs/infra/pkg/auth"
)

// Update ...
func (as *AuthStore) Update(a *auth.Auth) error {
	sql, args, err := squirrel.Update("auth").Set("blacklist", a.Blacklist).Where("id = ?", a.ID).Suffix("returning *").PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return err
	}

	row := as.Store.DB.QueryRowx(sql, args...)
	return row.StructScan(a)
}

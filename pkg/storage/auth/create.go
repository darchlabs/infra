package authstorage

import (
	"github.com/Masterminds/squirrel"
	"github.com/darchlabs/infra/pkg/auth"
)

// Create ...
func (as *AuthStore) Create(a *auth.Auth) error {
	sql, args, err := squirrel.
		Insert("auth").
		Columns("user_id", "token", "blacklist", "kind").
		Values(a.UserID, a.Token, a.Blacklist, a.Kind).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	row := as.Store.DB.QueryRowx(sql, args...)
	if err := row.StructScan(a); err != nil {
		return err
	}

	return nil
}

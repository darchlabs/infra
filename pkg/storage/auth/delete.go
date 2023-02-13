package authstorage

import (
	"database/sql"
	"time"

	"github.com/darchlabs/infra/pkg/auth"
)

// Delete ...
func (as *AuthStore) Delete(a *auth.Auth) error {
	row := as.Store.DB.QueryRowx("update auth set deleted_at = $1 where id = $2 returning *", time.Now(), a.ID)

	if err := row.StructScan(a); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	return nil
}

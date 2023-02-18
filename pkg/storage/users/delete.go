package userstorage

import (
	"database/sql"
	"time"

	"github.com/darchlabs/infra/pkg/users"
)

// Delete ...
func (us *UserStore) Delete(u *users.User) error {
	row := us.Store.DB.QueryRowx("update users set deleted_at = $1 where id = $2 returning *", time.Now(), u.ID)

	if err := row.StructScan(u); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	return nil
}

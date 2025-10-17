package repository

import (
	"log"
	er "mp/internal/errors"
	m "mp/internal/models"
)

func (r *RepositoryModule) SaveUser(user m.User) error {
	var us m.User
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE user_id=$1)`
	err := r.database.Db.Select(&us, query, user.UserId)
	if err != nil {
		log.Println(er.UserSelectError)
		return err
	}
	return nil
}

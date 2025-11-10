package repository

import (
	"errors"
	"log"
	er "mp/internal/errors"
	m "mp/internal/models"
	u "mp/internal/utils"

	"github.com/google/uuid"
)

func AdminInit(r *RepositoryModule) error {
	password, err := u.HashPassword("superadmin")
	if err != nil {
		return err
	}
	superAdmin := m.User{
		UserId:   uuid.New(),
		Login:    "superadmin",
		Password: password,
		Role:     "superadmin",
		Balance:  0,
	}
	var exist bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE user_role=$1)`
	err = r.database.Db.Get(&exist, query, "superadmin")
	if err != nil {
		log.Println("Superadmin create error")
		return err
	}
	if exist {
		return errors.New(er.AdminAlreadyExistErr)
	}
	err = r.SaveUser(superAdmin)
	if err != nil {
		log.Println("Superadmin save error", err)
		return err
	}
	return nil
}

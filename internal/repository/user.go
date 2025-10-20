package repository

import (
	"errors"
	"log"
	er "mp/internal/errors"
	m "mp/internal/models"
)

func (r *RepositoryModule) UserExist(user m.User) error {
	var exist bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE user_login=$1)`
	err := r.database.Db.Get(&exist, query, user.Login)
	if err != nil {
		log.Println(er.UserSelectError, err)
		return err
	}
	if exist {
		log.Println("User already exist")
		err = errors.New(er.UserAlreadyExistErr)
		return err
	}
	return nil
}

func (r *RepositoryModule) SaveUser(user m.User) error {
	err := r.UserExist(user)
	if err != nil {
		return err
	}
	query := `INSERT INTO users(user_id,user_login,user_password,user_role,user_balance)
	VALUES($1,$2,$3,$4,$5)`
	_, err = r.database.Db.Exec(query, user.UserId, user.Login, user.Password, user.Role, user.Balance)
	if err != nil {
		log.Println("User insert error:", err)
		return err
	}
	return nil
}

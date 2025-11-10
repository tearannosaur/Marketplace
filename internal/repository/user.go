package repository

import (
	"database/sql"
	"errors"
	"log"
	er "mp/internal/errors"
	m "mp/internal/models"

	"github.com/google/uuid"
)

func (r *RepositoryModule) UserExist(login string) error {
	var exist bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE user_login=$1)`
	err := r.database.Db.Get(&exist, query, login)
	if err != nil {
		log.Println(er.UserSelectError)
		return err
	}
	if exist {
		err = errors.New(er.UserAlreadyExistErr)
		return err
	}
	return nil
}

func (r *RepositoryModule) SaveUser(user m.User) error {
	err := r.UserExist(user.Login)
	if err != nil {
		return err
	}
	query := `INSERT INTO users(user_id,user_login,user_password,user_role,user_balance)
	VALUES($1,$2,$3,$4,$5)`
	_, err = r.database.Db.Exec(query, user.UserId, user.Login, user.Password, user.Role, user.Balance)
	if err != nil {
		log.Println("User insert error:")
		return err
	}
	return nil
}

func (r *RepositoryModule) GetUserLogin(login string) (m.User, error) {
	var user m.User
	query := `SELECT * FROM users WHERE user_login=$1`
	err := r.database.Db.Get(&user, query, login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println(er.UserDoesntExist)
			return m.User{}, err
		}
		return m.User{}, err
	}
	return user, nil
}

func (r *RepositoryModule) GetUserId(id uuid.UUID) (m.User, error) {
	var user m.User
	query := `SELECT * FROM users WHERE user_id=$1`
	err := r.database.Db.Get(&user, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println(er.UserDoesntExist)
			return m.User{}, err
		}
		return m.User{}, err
	}
	return user, nil
}

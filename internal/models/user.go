package models

import (
	"errors"
	"log"
	"mp/internal/utils"

	"github.com/google/uuid"
)

type User struct {
	User_id  uuid.UUID `json:"user_id"`
	Login    string    `json:"user_login"`
	Password string    `json:"user_password"`
	Role     string    `json:"user_role"`
	Balance  float64   `json:"user_balance"`
}

type UserResponse struct {
	User_id uuid.UUID `json:"user_id"`
	Login   string    `json:"user_login"`
}

type UserRequest struct {
	Login    string `json:"user_login"`
	Password string `json:"user_password"`
}

// CreateUser создаёт нового пользователя из данных запроса.
// Возвращает ошибку, если не удалось сгенерировать хэш пароля.
func CreateUser(u UserRequest) (User, error) {
	password, err := utils.HashPassword(u.Password)
	if err != nil {
		log.Println("Ошибка генерации хэша", err)
		return User{}, err
	}
	if u.Password == "" || u.Login == "" {
		log.Println("Неккоректные данные пользователя", u.Login)
		err = errors.New("неккоректные данные пользователя")
		return User{}, err
	}
	return User{
		User_id:  uuid.New(),
		Login:    u.Login,
		Password: password,
		Role:     "user",
		Balance:  0,
	}, nil
}

// Withdraw — снимает деньги с баланса пользователя.
// Возвращает ошибку, если на счёте недостаточно средств.
func (u *User) Withdraw(amount float64) error {
	if u.Balance-amount < 0 {
		return errors.New("недостаточно средств на счету")
	}
	u.Balance -= amount
	return nil
}

// Deposit — пополняет баланс пользователя на указанную сумму.
// Возвращает ошибку, если сумма некорректна.
func (u *User) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма должна быть положительной")
	}
	u.Balance += amount
	return nil
}

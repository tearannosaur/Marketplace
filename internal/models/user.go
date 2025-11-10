package models

import (
	"errors"
	"log"
	er "mp/internal/errors"
	"mp/internal/utils"

	
	"github.com/google/uuid"
)

type User struct {
	UserId   uuid.UUID `json:"user_id" db:"user_id"`
	Login    string    `json:"user_login" db:"user_login"`
	Password string    `json:"user_password" db:"user_password"`
	Role     string    `json:"user_role" db:"user_role"`
	Balance  float64   `json:"user_balance" db:"user_balance"`
}

type UserResponse struct {
	UserId uuid.UUID `json:"user_id"`
	Login  string    `json:"user_login"`
}

type UserRequest struct {
	Login    string `json:"user_login"`
	Password string `json:"user_password"`
}

// ValidateUserData проверяет данные Json запроса.
// Возвращает ошибку если данные не корректны.
func ValidateUserData(u UserRequest) error {
	if u.Password == "" || u.Login == "" {
		err := errors.New(er.IncorrectJsonBody)
		return err
	}
	return nil
}

// CreateUser создаёт нового пользователя из данных запроса.
// Возвращает ошибку, если не удалось сгенерировать хэш пароля.
func CreateUser(u UserRequest) (User, error) {
	err := ValidateUserData(u)
	if err != nil {
		return User{}, err
	}
	password, err := utils.HashPassword(u.Password)
	if err != nil {
		log.Println(er.HashGenerateErr)
		return User{}, err
	}

	return User{
		UserId:   uuid.New(),
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
		return errors.New(er.InsufficientFundsErr)
	}
	u.Balance -= amount
	return nil
}

// Deposit — пополняет баланс пользователя на указанную сумму.
// Возвращает ошибку, если сумма некорректна.
func (u *User) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New(er.InvalidDepositAmountErr)
	}
	u.Balance += amount
	return nil
}

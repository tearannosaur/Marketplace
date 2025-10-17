package utils

import (
	"errors"

	er "mp/internal/errors"

	b "golang.org/x/crypto/bcrypt"
)

// HashPassword генерирует хэш из пароля.
// Возвращает хэш и ошибку, если генерация не удалась.
func HashPassword(password string) (string, error) {
	if password == "" {
		return password, errors.New(er.EmptyPasswordError)
	}
	p := []byte(password)
	hash, err := b.GenerateFromPassword(p, b.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// VerifyPassword проверяет, соответствует ли пароль хэшу.
func VerifyPassword(hash string, password string) error {
	h := []byte(hash)
	p := []byte(password)
	if err := b.CompareHashAndPassword(h, p); err != nil {
		return err
	}
	return nil
}

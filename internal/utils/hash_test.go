package utils

import (
	"fmt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	passwords := []string{
		"password1",
		"password",
		"fdsm,l;sd",
		"fdsm !,l;sd",
	}
	for _, password := range passwords {
		t.Run(fmt.Sprintf("password:%s", password), func(t *testing.T) {
			hash, err := HashPassword(password)
			if err != nil {
				t.Fatalf("Hash password error:%v", err)
			}
			err = VerifyPassword(hash, password)
			if err != nil {
				t.Error("Verify password error:", err)
			}
			err = VerifyPassword(hash, password+"213jm1k")
			if err == nil {
				t.Error("VerifyPassword passed for wrong password, want error")
			}

		})

	}
	t.Run("empty password", func(t *testing.T) {
		empty := ""
		_, err := HashPassword(empty)
		if err == nil {
			t.Fatalf("Empty password,want error.")
		}
	})
}

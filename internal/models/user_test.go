package models

import (
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {
	users := []UserRequest{{
		Login:    "login",
		Password: "password",
	}, {
		Login:    "logiabcvn",
		Password: "as",
	},
	}
	for _, user := range users {
		t.Run(fmt.Sprintf("user:%v", user), func(t *testing.T) {
			us, err := CreateUser(user)
			if err != nil {
				t.Fatalf("Cannot create new user error:%v", us)
			}
			if us.Role != "user" {
				t.Errorf("Wrong user role,need:user")
			}
			if us.Balance != 0 {
				t.Errorf("Wrong user balance,need:0")
			}
			//проверять хэш
		})

	}
	invalidUsers := []UserRequest{
		{Login: "",
			Password: "passwordsa"},
		{Login: "sda",
			Password: ""},
		{
			Login:    "",
			Password: "",
		},
	}
	for _, us := range invalidUsers {
		t.Run(fmt.Sprintf("user:%v", us), func(t *testing.T) {
			_, err := CreateUser(us)
			if us.Password == "" || us.Login == "" {
				if err == nil {
					t.Fatalf("Empty user data,want error")
				}
			}

		})
	}

}

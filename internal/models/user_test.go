package models

import (
	"fmt"
	"mp/internal/utils"
	"testing"
)

func TestCreateUser(t *testing.T) {
	users := []UserRequest{
		{Login: "login", Password: "password"},
		{Login: "logiabcvn", Password: "as"},
	}
	for _, user := range users {
		t.Run(fmt.Sprintf("user:%v", user), func(t *testing.T) {
			us, err := CreateUser(user)
			if err != nil {
				t.Fatalf("Cannot create new user%v ,error:%v", us, err)
			}
			if us.Role != "user" {
				t.Errorf("Expected role 'user', got '%s' for login %s", us.Role, user.Login)
			}
			if us.Balance != 0 {
				t.Errorf("Wrong user balance,need:0")
			}
			err = utils.VerifyPassword(us.Password, user.Password)
			if err != nil {
				t.Errorf("Password verification failed for user %s", user.Login)
			}
		})

	}
	invalidUsers := []UserRequest{
		{Login: "", Password: "passwordsa"},
		{Login: "sda", Password: ""},
		{Login: "", Password: ""},
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

func TestWithdraw(t *testing.T) {
	users := []User{
		{Balance: 10},
		{Balance: 100},
	}
	for _, v := range users {
		t.Run(fmt.Sprintf("Balance:%v", v.Balance), func(t *testing.T) {
			before := v.Balance - 10
			err := v.Withdraw(10)
			if err != nil {
				t.Errorf("Withdraw error:%v", err)
			}

			if before != v.Balance {
				t.Errorf("Expected balance,want:%v,got:%v", before, v.Balance)
			}
		})
	}
	invalidUsers := []User{
		{Balance: 0},
		{Balance: -100},
	}
	for _, v := range invalidUsers {
		t.Run(fmt.Sprintf("Balance:%v", v.Balance), func(t *testing.T) {
			err := v.Withdraw(10)
			if err == nil {
				t.Errorf("Not enough money,want error")
			}
			if v.Balance != 0 && v.Balance != -100 {
				t.Errorf("Balance changed incorrectly: %v", v.Balance)
			}
		})
	}
}

func TestDeposit(t *testing.T) {
	users := []User{
		{Balance: 10},
		{Balance: 100},
		{Balance: -100},
	}
	for _, v := range users {
		t.Run(fmt.Sprintf("Balance:%v", v.Balance), func(t *testing.T) {
			need := v.Balance + 100
			err := v.Deposit(100)
			if err != nil {
				t.Errorf("Deposit error:%v", err)
			}
			if need != v.Balance {
				t.Errorf("Balance changed incorrectly,want:%v,got:%v", need, v.Balance)
			}
		})
	}
}

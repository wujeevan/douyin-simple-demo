package repository

import (
	"testing"
)

func TestQueryUserById(t *testing.T) {
	QueryUserById(1)
}

func TestQueryUserByUsername(t *testing.T) {
	QueryUserByUsername("adminssss")
}

func TestUpdateUserToken(t *testing.T) {
	UpdateUserToken(1, "aiewdasdjdsd")
}

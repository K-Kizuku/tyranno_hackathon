package logic_test

import (
	"testing"
	"tyranno/backend/utils/logic"
)

func TestHashPassword(t *testing.T) {
	password := "testPassword123"
	hashed, err := logic.HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if len(hashed) == 0 {
		t.Error("Expected hashed password to be non-empty")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "testPassword123"
	hashed, _ := logic.HashPassword(password)

	if !logic.CheckPasswordHash(password, hashed) {
		t.Error("Expected original password and hashed password to match")
	}

	wrongPassword := "wrongPassword"
	if logic.CheckPasswordHash(wrongPassword, hashed) {
		t.Error("Expected wrong password and hashed password not to match")
	}
}

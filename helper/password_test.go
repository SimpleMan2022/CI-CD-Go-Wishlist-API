package helper

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := "admin123"
	hashPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("Error Hashing Password: %v", err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)); err != nil {
		t.Errorf("Hashed Password not valid: %v", err)
	}

}

func TestVerifyPassword(t *testing.T) {
	password := "admin123"

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err := VerifyPassword(password, string(hashPassword)); err != nil {
		t.Errorf("Failed to verify password: %v", err)
	}
}

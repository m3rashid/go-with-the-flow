package auth

import "testing"

func TestHashPassword(t *testing.T) {
	password := "password"
	hashedPassword := HashPassword(password)
	if hashedPassword == password {
		t.Error("HashPassword() failed")
	}
}

func TestVerifyPassword(t *testing.T) {
	password := "password"
	hashedPassword := HashPassword(password)
	if !VerifyPassword(hashedPassword, password) {
		t.Error("VerifyPassword() failed")
	}
}

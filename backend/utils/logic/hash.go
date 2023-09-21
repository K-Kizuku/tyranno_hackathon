package logic

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword は与えられたパスワードをハッシュ化します。
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash はハッシュ化されたパスワードと元のパスワードがマッチするかを確認します。
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

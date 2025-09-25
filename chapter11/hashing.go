package chapter11

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Hashing() {
	password := "Password"
	hash, _ := hashPassword(password)

	fmt.Println("password:", password)
	fmt.Println("hash:", hash)

	match := checkPassword(password, hash)
	fmt.Println(match)
}

func hashPassword(password string) (string, error) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), nil
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

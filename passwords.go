package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Password(password string) {
	hash, err := HashPassword(password)
	if err != nil {
		panic(err)
	}
	fmt.Println("password = ", password)
	fmt.Println(hash)

	check := CheckPasswordHash(password, hash)
	fmt.Println(check)
}

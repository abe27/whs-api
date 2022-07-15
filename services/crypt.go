package services

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Println(err)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
	fmt.Println(err)
	h, _ := HashPassword(password)
	fmt.Println("PASSWORD: " + password)
	fmt.Println("HASH: " + hash)
	fmt.Println("HASH PASSWORD: " + h)
	fmt.Println("+++++++|+++++++")
	return err == nil
}

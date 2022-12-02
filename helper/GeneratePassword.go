package helper

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashAndSalt(strPassword []byte) string {
	password, err := bcrypt.GenerateFromPassword(strPassword, 10)
	PanicIfError(err)

	return string(password)
}

func ComparePassword(hashPassword string, plainPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

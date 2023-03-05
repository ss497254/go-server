package utils

import (
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type T interface{}

func CheckErr(x T, err error) T {
	if err != nil {
		fmt.Println("Closing from CheckErr")
		log.Fatal(err)
	}

	return x
}

func JSON_Stringify(v any) string {
	res, err := json.Marshal(v)

	if err != nil {
		fmt.Println("Unable to parse JSON")
		return "{}"
	}

	return string(res)
}

func ComparePasswords(hashedPassword string, plainPassword []byte) bool {
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	return err == nil
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	if err != nil {
		return ""
	}

	return string(hash)
}

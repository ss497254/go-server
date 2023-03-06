package lib

import "golang.org/x/crypto/bcrypt"

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

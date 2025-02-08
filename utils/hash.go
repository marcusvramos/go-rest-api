package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bcryptPassword), err
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
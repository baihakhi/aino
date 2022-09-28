package crypt

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (string, error) {

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(encryptedPassword), nil
}

func ComparePassword(password, passwordDB string) error {

	err := bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

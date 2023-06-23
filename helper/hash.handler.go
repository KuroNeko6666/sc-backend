package helper

import "golang.org/x/crypto/bcrypt"

func CompareHash(hash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func GenerateHash(password *string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), 14)
	if err != nil {
		return err
	}
	*password = string(bytes)
	return nil
}

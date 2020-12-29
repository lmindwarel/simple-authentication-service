package utils

import (
	"github.com/dchest/uniuri"
	"golang.org/x/crypto/bcrypt"
)

// HashFromString generate hash from given string
func HashFromString(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hash), err
}

// HashCorrespond check if given passord correspond to the hash
func HashCorrespond(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}

// RandomHash return random string with given length
func RandomHash(length int) string {
	return uniuri.NewLen(length)
}

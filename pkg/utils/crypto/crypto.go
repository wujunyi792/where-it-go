package crypto

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Md5Crypto(password string) string {
	data := []byte(password)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func Md5CryptoWithSalt(password string, salt string) string {
	data := []byte(password + salt)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func PasswordGen(password string, salt string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(Md5CryptoWithSalt(password, salt)), bcrypt.DefaultCost)
	return string(hash)
}

func PasswordCompare(passwordInput string, correctPassword string, salt string) bool {
	return bcrypt.CompareHashAndPassword([]byte(correctPassword), []byte(Md5CryptoWithSalt(passwordInput, salt))) == nil
}

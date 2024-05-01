package passwords

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
)

const SALT = "passfish"

func secureRandomInt(max int) int {
	nextInt, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		log.Fatal(err)
	}
	return int(nextInt.Int64())
}

func GeneratePassword(length int) string {
	const charset = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_-+=~ `
	const charsetLength = len(charset)

	password := ""
	var nextInt int
	for i := 0; i < length; i++ {
		nextInt = secureRandomInt(charsetLength)
		password += string(charset[nextInt])
	}
	return password
}

type Login struct {
	Login    string "json:\"login\""
	Username string "json:\"username\""
	Password string "json:\"password\""
}

func NewLogin(login string, username string, password string) Login {
	return Login{login, username, password}
}

func (login *Login) String() string {
	return fmt.Sprintf("Login{Login: %s, Username: %s, Password: %s}", login.Login, login.Username, "XXXXXXXXXX")
}

func (login *Login) Encrypt() string {
	data := []byte(SALT + login.Login + login.Username + login.Password)
	str := fmt.Sprintf("%x", sha256.Sum256(data))[:32]
	fmt.Println("Encrypted Login: ", str)
	return str
}

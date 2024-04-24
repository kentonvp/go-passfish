package passwords

import (
	"fmt"
	"math/rand"
)

func GeneratePassword() string {
	return fmt.Sprintf("password%d", rand.Intn(50000))
}

type Login struct {
	Login    string "json:\"login\""
	Username string "json:\"username\""
	Password string "json:\"password\""
}

func (Login) String() string {
	return fmt.Sprintf
}

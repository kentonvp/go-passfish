package models

import (
	"fmt"
	"passfish/internal/passwords"
)

type Credentials struct {
	Title    string `json:"title"`
	Username string `json:"username"`
	Password string `json:"password"` // Password should never be plain text.
}

func (creds *Credentials) String() string {
	return fmt.Sprintf("Credentials{Title: %s, Username: %s, Password: %s}", creds.Title, creds.Username, creds.Password)
}

func (creds *Credentials) DecryptPassword(passphrase string) string {
	return passwords.Decrypt(creds.Password, passphrase)
}

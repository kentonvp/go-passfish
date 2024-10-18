package models

import (
	"fmt"
	"passfish/internal/passwords"
	"time"
)

type BaseCredentials struct {
	Title    string `json:"title"`
	Username string `json:"username"`
	Password string `json:"password"` // Password should never be plain text.
}

type Credentials struct {
	Base         BaseCredentials
	LastAccessed time.Time `json:"last_accessed"`
}

func (creds *Credentials) String() string {
	return fmt.Sprintf("Credentials{Title: %s, Username: %s, Password: [REDACTED], LastAccessed: %s}", creds.Base.Title, creds.Base.Username, creds.LastAccessed.String())
}

func (creds *Credentials) DecryptPassword(passphrase string) string {
	return passwords.Decrypt(creds.Base.Password, passphrase)
}

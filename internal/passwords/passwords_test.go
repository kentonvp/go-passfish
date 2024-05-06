package passwords_test

import (
	"passfish/internal/passwords"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	plaintext := "secret"
	passphrase := "abracadabra"

	encrypted := passwords.Encrypt(plaintext, passphrase)
	decrypted := passwords.Decrypt(encrypted, passphrase)

	if plaintext != decrypted {
		t.Errorf("Expected %s, got %s", plaintext, decrypted)
	}
}

package passwords

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"log"
	"math/big"
)

func secureRandomInt(max int) int {
	nextInt, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		log.Fatal(err)
	}
	return int(nextInt.Int64())
}

const charset string = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_-+=~ `
const charsetLength int = len(charset)

func New(length int) string {
	var (
		nextInt  int
		password string = ""
	)

	for i := 0; i < length; i++ {
		nextInt = secureRandomInt(charsetLength)
		password += string(charset[nextInt])
	}

	return password
}

func sha256Hash(input string) []byte {
	plaintext := []byte(input)
	h := sha256.Sum256(plaintext)
	return h[:]
}

func Encrypt(plaintext string, passphrase string) string {
	key := sha256Hash(passphrase)

	// Code taken from: https://dev.to/breda/secret-key-encryption-with-go-using-aes-316d
	aes, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal("error creating cipher", err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		log.Fatal("error creating gcm", err)
	}

	// We need a 12-byte nonce for GCM (modifiable if you use cipher.NewGCMWithNonceSize())
	// A nonce should always be randomly generated for every encryption.
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		log.Fatal(err)
	}

	// ciphertext here is actually nonce+ciphertext
	// So that when we decrypt, just knowing the nonce size
	// is enough to separate it from the ciphertext.
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	return string(ciphertext)
}

func Decrypt(ciphertext string, passphrase string) string {
	key := sha256Hash(passphrase)

	aes, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal("error creating cipher", err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		log.Fatal("error creating gcm", err)
	}

	// Since we know the ciphertext is actually nonce+ciphertext
	// And len(nonce) == NonceSize(). We can separate the two.
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		log.Fatal("error opening gcm", err)
	}

	return string(plaintext)
}

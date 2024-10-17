package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

func hashTo32Bytes(input string) []byte {
	hash := sha256.Sum256([]byte(input))
	return hash[:]
}

// Encrypt Key using AES encryption
func EncryptKey(key, secretKeyStr string) (string, error) {
	hashedSecretKey := hashTo32Bytes(secretKeyStr)
	// Create a new AES cipher using the secret key
	block, err := aes.NewCipher(hashedSecretKey)
	if err != nil {
		return "", err
	}

	// GCM provides authenticated encryption with associated data (AEAD)
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Create a nonce. GCM standard recommends a 12-byte nonce size.
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt the key
	ciphertext := aesGCM.Seal(nonce, nonce, []byte(key), nil)

	// Return the encrypted text as a base64-encoded string
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt Key using AES encryption
func DecryptKey(encryptedKey, secretKeyStr string) (string, error) {
	hashedSecretKey := hashTo32Bytes(secretKeyStr)

	// Decode the base64-encoded string
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedKey)
	if err != nil {
		return "", err
	}

	// Create a new AES cipher using the secret key
	block, err := aes.NewCipher(hashedSecretKey)
	if err != nil {
		return "", err
	}

	// GCM provides authenticated encryption with associated data (AEAD)
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Extract the nonce from the ciphertext
	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt the key
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

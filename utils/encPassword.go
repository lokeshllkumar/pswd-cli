package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func EncryptPassword(password string) (string, error) {
	block, err := aes.NewCipher(Key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	encPassword := aesGCM.Seal(nonce, nonce, []byte(password), nil)
	return base64.StdEncoding.EncodeToString(encPassword), nil
}

func DecryptPassword(encPassword string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encPassword)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(Key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("ciphertext is too short")
	}

	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	decPassword, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		fmt.Printf("Error decrypting password\n")
		return "", err
	}

	return string(decPassword), nil
}

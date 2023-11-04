package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

func AESencrypt(key string, textToEncrypt string) (string, error) {
	textInBytes := []byte(textToEncrypt)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	cipheredText := make([]byte, aes.BlockSize+len(textToEncrypt))
	iv := cipheredText[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipheredText[aes.BlockSize:], textInBytes)

	encryptedText := base64.URLEncoding.EncodeToString(cipheredText)
	return encryptedText, nil
}

func AESdecrypt(key string, encryptedText string) (string, error) {
	keyInBytes := []byte(key)
	cipheredText, _ := base64.URLEncoding.DecodeString(encryptedText)

	block, err := aes.NewCipher(keyInBytes)
	if err != nil {
		return "", err
	}

	if len(cipheredText) < aes.BlockSize {
		return "", errors.New("ERROR: couldn't decrypt, the size of the encrypted text must be at least 16 bytes")
	}
	iv := cipheredText[:aes.BlockSize]
	cipheredText = cipheredText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(cipheredText, cipheredText)

	return string(cipheredText), nil
}

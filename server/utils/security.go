package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

func GenerateAuthToken(username string) (string, error) {
	tokenBytes := make([]byte, 32)

	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	token := base64.StdEncoding.EncodeToString(tokenBytes)
	token = username + "-" + token // ensure it does not colide with other auth tokens

	return token, nil
}

func GenerateRandomSalt() ([]byte, error) {
	var salt = make([]byte, 16)

	_, err := rand.Read(salt[:])

	return salt, err
}

func HashPasswordSalt(password string, salt []byte) (string, error) {
	var passwordBytes = []byte(password)
	var sha512Hasher = sha512.New()

	passwordBytes = append(passwordBytes, salt...)
	_, err := sha512Hasher.Write(passwordBytes)

	if err != nil {
		return "", err
	}

	var hashedPasswordBytes = sha512Hasher.Sum(nil)
	var hashedPasswordHex = hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex, err
}

func ValidatePassword(hashedPassword string, currPassword string, salt []byte) (bool, error) {
	currPasswordHash, err := HashPasswordSalt(currPassword, salt)

	if err != nil {
		return false, err
	}

	return hashedPassword == currPasswordHash, nil
}

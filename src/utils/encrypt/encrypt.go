package encrypt

import (
	"encoding/hex"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

func PasswordsMatch(password, hash, encryptedPassword string) bool {
	salt, err := decodeHashFromString(hash)
	if err != nil {
		return false
	}

	hashedPassword, err := decodeHashFromString(encryptedPassword)
	if err != nil {
		return false
	}

	bytes := applySaltToPassword(password, salt)
	err = bcrypt.CompareHashAndPassword(hashedPassword, bytes)
	if err != nil {
		return false
	}

	return true
}

func applySaltToPassword(password string, salt []byte) []byte {
	return append([]byte(password), salt[:]...)
}

func decodeHashFromString(hash string) ([]byte, error) {
	decoded, err := hex.DecodeString(hash)
	if err != nil {
		log.Errorf("There was an error decoding the password: %v", err.Error())
		return nil, err
	}

	return decoded, nil
}

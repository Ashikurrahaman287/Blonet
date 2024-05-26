package utils

import (
	"crypto/ecdsa"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func LoadECDSAKey(filepath string) (*ecdsa.PrivateKey, error) {
	keyJson, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	password := "your-keystore-password"
	key, err := keystore.DecryptKey(keyJson, password)
	if err != nil {
		return nil, err
	}

	return key.PrivateKey, nil
}

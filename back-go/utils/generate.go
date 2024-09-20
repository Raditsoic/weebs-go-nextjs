package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateUniqueID() string {
	n, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return ""
	}
	return n.String()
}

package random

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"strconv"
)

func String(length int) (string, error) {
	b := make([]byte, 2*length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func StringWithSource(length int, source int) (string, error) {
	nonce := "qrAMQmTNzpcArj5SP7i3sVOuKpdrDgT0"

	i := source << 3
	n, err := strconv.Itoa(i)
	if err != nil {
		return "", err
	}
	dat := []byte(n + nonce)
	return hex.EncodeToString(dat)
}

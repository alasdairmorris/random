package main

import (
	"crypto/rand"
	"math/big"
)

func randomString(charset []rune, length int) string {
	b := make([]rune, length)
	for i := range b {
		pos, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		b[i] = charset[pos.Uint64()]
	}
	return string(b)
}

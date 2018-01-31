package main

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"os"
	"testing"
	"time"
)

var pinLength = uint8(4)
var availableChars = []byte("0123456789")

func BenchmarkBrutforcePIN(b *testing.B) {
	hash := getHash()
	key := getKey()
	keySize := uint8(len(key))

	for n := 0; n < b.N; n++ {
		res, err := BrutSha256(hash, availableChars, keySize)
		if err != nil {
			panic(err)
		}

		if res != key {
			panic("wrong key")
		}
	}
}

func getHash() string {
	hash := os.Getenv("HASH_FOR_BRUTFORCE")
	if len(hash) > 0 {
		return hash
	}

	pin := generatePINWithLen(int(pinLength), availableChars)

	hasher := sha256.New()
	hasher.Write([]byte(pin))
	hashBin := hasher.Sum(nil)
	hash = hex.EncodeToString(hashBin)

	os.Setenv("HASH_FOR_BRUTFORCE", hash)
	os.Setenv("KEY_FOR_BRUTFORCE", pin)

	return hash
}

func getKey() string {
	return os.Getenv("KEY_FOR_BRUTFORCE")
}

func generatePINWithLen(length int, availableChars []byte) string {
	letters := availableChars

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	b := make([]byte, length)
	for i := range b {
		b[i] = letters[r1.Intn(len(letters))]
	}

	return string(b)
}

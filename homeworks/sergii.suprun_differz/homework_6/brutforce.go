package main

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/MastersAcademy/go-course-2017/homeworks/sergii.suprun_differz/homework_6/brut"
)

// BrutSha256 returns source (we would call it `key`) of hashsum sha256
// params:
// hash - sha256 hash, represented as HEX string
// availableChars - chars that could be used in key (if it PIN than available chars would be - 0,1,2,3,4,5,6,7,8,9)
// length - length of key
// e.g.:
// BrutSha256("888df25ae35772424a560c7152a1de794440e0ea5cfee62828333a456a506e05", []byte("1234567890"), 4) should return "9999", nil
func BrutSha256(hash string, availableChars []byte, length uint8) (string, error) {
	return brut.NewSha256(hash, availableChars, length).Run()
}

func main() {
	key := "9999"
	hasher := sha256.New()
	hasher.Write([]byte(key))
	hash := hasher.Sum(nil)
	hashHex := hex.EncodeToString(hash)

	result, err := BrutSha256(hashHex, []byte("1234567890"), 4)
	if err != nil {
		panic(err)
	}

	if result != key {
		panic("hash not decrypted")
	}
}

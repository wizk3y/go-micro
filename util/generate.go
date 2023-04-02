package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	ASCIILowercase = "abcdefghijklmnopqrstuvwxyz"
	ASCIIUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ASCIILetter    = ASCIILowercase + ASCIIUppercase
	Digits         = "0123456789"
	HexDigits      = Digits + "abcdef" + "ABCDEF"
	OctDigits      = "01234567"
)

func GenerateString(n int, letters string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}

func GenerateStringWithASCIILetter(n int) string {
	return GenerateString(n, ASCIILetter)
}

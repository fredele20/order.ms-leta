package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generate a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomName generates a random owner name
func RandomName() string {
	return RandomString(6)
}

// RandomDescription generates a random description for a product
func RandomDescription() string {
	return RandomString(12)
}

// RandomAmount generates a random amount of money
func RandomPrice() int64 {
	return RandomInt(0, 1000)
}

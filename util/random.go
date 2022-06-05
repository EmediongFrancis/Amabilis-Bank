package util

import (
	"math/rand"
	"strings"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Initialize random number generator.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Returns a random int in the range [min, max].
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Returns a random string of length n.
func RandomString(n int) string {
	var sb strings.Builder
	k := len(letters)

	for i := 0; i < n; i++ {
		c := letters[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// Generate random account owner.
func RandomOwner() string {
	return RandomString(7)
}

// Generate random currency.
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "GBP", "JPY", "CAD", "NGN", "GHC"}
	N := len(currencies)
	return currencies[rand.Intn(N)]
}

// Generate random account balance.
func RandomBalance() int64 {
	return RandomInt(1, 100)
}

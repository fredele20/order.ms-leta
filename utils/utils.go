package utils

import (
	"math/rand"
	"strings"
	"time"

	"github.com/oklog/ulid"
)

// This function is used to generate a unique id
func GenerateId() string {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().Unix())), 0)
	return strings.ToLower(ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String())
}

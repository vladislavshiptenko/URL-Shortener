package random

import (
	"math/rand"
	"time"
)

func NewRandomString(aliasLength int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	var result string
	for i := 0; i < aliasLength; i++ {
		result += string(rune(97 + rnd.Intn(26)))
	}

	return result
}

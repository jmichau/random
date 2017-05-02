// Package random generates fixed length random strings.
// Code from this package is originally created by icza (http://stackoverflow.com/users/1705598/icza),
// posted here: http://stackoverflow.com/a/31832326
// and licensed under cc by-sa 3.0
package random

import (
	"math/rand"
	"time"
)

var letterBytes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var src = rand.NewSource(time.Now().UnixNano())

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// Generate a random string of a fixed length from [a-zA-Z] range.
func Generate(length int) string {
	b := make([]byte, length)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = byte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

package utils

import (
	"math/rand"
)

const (
	characters = "[]()/.,*&"
	charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func GetRandChars(l int) string {
	s := make([]byte, l)
	for i := range s {
		s[i] = charset[rand.Intn(len(charset))]
	}
	return string(s)
}

func GetRandCharacters(l int) string {
	s := make([]byte, l)
	for i := range s {
		s[i] = characters[rand.Intn(len(characters))]
	}
	return string(s)
}

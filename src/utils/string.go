package utils

import "math/rand"

type String struct {
	//
}

func (i String) Random(length int) string {
	var availableLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	letters := make([]rune, length)
	for x := range letters {
		letters[x] = availableLetters[rand.Intn(len(availableLetters))]
	}

	return string(letters)
}

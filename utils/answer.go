package utils

import "strings"

func CleanUpAnswer(word string) string {
	c := strings.ToLower(word)
	c = strings.TrimSpace(c)
	return c
}
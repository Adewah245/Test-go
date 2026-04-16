package main

import (
	"strings"
)

func atoan(s string) string {
	word := strings.Fields(s)

	for i := 0; i < len(word); i++ {
		vowel := "aeiouhAEIOUH"
		if word[i] == "a" && strings.Contains(vowel, string(word[i+1][0])) {
			word[i] = "an"
		}
		if word[i] == "an" && !strings.Contains(vowel, string(word[i+1][0])) {
			word[i] = "a"
		}
		if word[i] == "A" && !strings.Contains(vowel, string(word[i+1][0])) {
			word[i] = "An"
		}
		if word[i] == "An" && !strings.Contains(vowel, string(word[i+1][0])) {
			word[i] = "A"
		}
	}
	return strings.Join(word, " ")
}

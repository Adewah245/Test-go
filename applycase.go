package main

import (
	"strings"
)

func applycase(s string) string {
	word := strings.Fields(s)

	for i := 0; i < len(word); i++ {
		if word[i] == "(up)" {
			if i > 0 {
				word[i-1] = strings.ToUpper(word[i-1])
			}
			word = append(word[:i], word[i+1:]...)
			i--
		}
		if word[i] == "(low)" {
			if i > 0 {
				word[i-1] = strings.ToLower(word[i-1])
			}
			word = append(word[:i], word[i+1:]...)
			i--
		}
		if word[i] == "(cap)" {
			if i > 0 && len(word[i-1]) > 0 {
				word[i-1] = strings.ToUpper(word[i-1][:1]) + word[i-1][1:]
			}
			word = append(word[:i], word[i+1:]...)
			i--
		}

	}
	return strings.Join(word, " ")
}

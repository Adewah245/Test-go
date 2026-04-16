package main

import (
	"strconv"
	"strings"
)

func convert(s string) string {
	word := strings.Fields(s)

	for i := 0; i < len(word); i++ {
		switch word[i] {
		case "(hex)":
			data, _ := strconv.ParseInt(word[i-1], 16, 64)
			word[i-1] = strconv.FormatInt(data, 10)
			word = append(word[:i], word[i+1:]...)
			i--
		case "(bin)":
			data, _ := strconv.ParseInt(word[i-1], 2, 64)
			word[i-1] = strconv.FormatInt(data, 10)
			word = append(word[:i], word[i+1:]...)
			i--
		}
	}
	return strings.Join(word, " ")
}

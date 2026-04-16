package main

import (
	"strings"
)

func fixpuct(s string) string {
	s = strings.ReplaceAll(s, " ,", ",")
	s = strings.ReplaceAll(s, ",", ", ")
	s = strings.ReplaceAll(s, " !", "!")
	s = strings.ReplaceAll(s, " .", ".")
	s = strings.ReplaceAll(s, " ?", "?")
	return s
}

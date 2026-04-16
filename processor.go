package main

import "strings"

func processor(s string) string {
	line := strings.Split(s, "\n")
	for i := 0; i < len(line); i++ {
		line[i] = applycase(line[i])
		line[i] = atoan(line[i])
		line[i] = convert(line[i])
		line[i] = fixpuct(line[i])
		line[i] = lastnword(line[i])
		line[i] = fixquote(line[i])
	}
	return strings.Join(line, "\n")
}

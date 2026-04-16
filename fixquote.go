package main

import "regexp"

func fixquote(s string) string {
	single := regexp.MustCompile(`'\s+(.*?)\s+'`)
	s = single.ReplaceAllString(s, `'$1'`)
	second := regexp.MustCompile(`"\s+(.*?)\s+"`)
	s = second.ReplaceAllString(s, `"$1"`)

	return s
}

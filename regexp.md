```md
# Regular expression = (regexp)
***meaning**
A way to describe patterns in text.

**symbols  and their meaning**
1. dot (.) simply means every character
2. * => means zero or more
3. + => means one or more
4. ? => 0 or 1 time
5. [] => character set. (one of these)
6. () => group parts
7. | => OR
8. {n} => exact repeat
9. {n,m} => range repeat
10. ^  => start of text
11. $ => end of text
12. \ => escape / special meaning
13. \d => digit
14. \s => space
15. \b => boundary
```
```txt
mostly used in regexp:
[] () \s . ? + $ *
```
# processor.go
```go
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
```
# fixquote.go
```go
package main

import "regexp"

func fixquote(s string) string {
	single := regexp.MustCompile(`'\s+(.*?)\s+'`)
	s = single.ReplaceAllString(s, `'$1'`)
	second := regexp.MustCompile(`"\s+(.*?)\s+"`)
	s = second.ReplaceAllString(s, `"$1"`)

	return s
}
```



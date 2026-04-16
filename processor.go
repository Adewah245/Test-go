package main

func processor(s string) string {
	s = applycase(s)
	s = lastnword(s)
	s = convert(s)
	s = atoan(s)
	s = fixpuct(s)
	return s
}

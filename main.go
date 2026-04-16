package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . sample.txt result.txt")
		os.Exit(1)
	}

	data, _ := os.ReadFile(os.Args[1])

	result := processor(string(data))

	os.WriteFile(os.Args[2], []byte(result), 0644)
	fmt.Println(result)
}

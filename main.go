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

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Erro Reading File", err)
		os.Exit(1)
	}

	result := processor(string(data))
	err = os.WriteFile(os.Args[2], []byte(result), 0644)
	fmt.Println(result)
}

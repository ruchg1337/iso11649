package main

import (
	"fmt"
	"os"

	iso11649 "github.com/ruchg1337/iso11649"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <input>")
		os.Exit(1)
	}

	input := os.Args[1]
	result := iso11649.GenerateRfReference(input)
	fmt.Println(result)
}

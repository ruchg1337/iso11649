package main

import (
	"fmt"
	"os"

	iso11649 "github.com/ruchg1337/iso11649"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: iso11649 <input>")
		os.Exit(1)
	}

	input := os.Args[1]
	reference, err := iso11649.GenerateReference(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(reference)
}

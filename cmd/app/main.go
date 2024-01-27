package main

import (
	"calculator-parser/internal"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage ./calc 5 + 5 - 1")
	}

	input := strings.Join(os.Args[1:], "")
	parser := internal.NewParser()

	result, err := parser.Parse(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %d", result)
}

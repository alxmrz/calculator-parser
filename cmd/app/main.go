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
	calculator := internal.NewCalculator(internal.NewParser())

	result := calculator.Calculate(input)

	fmt.Printf("Result: %d", result)
}

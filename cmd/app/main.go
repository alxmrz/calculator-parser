package main

import (
	"calculator-parser/internal"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage ./calc \"5 + 5 - 1\"")
	}

	input := os.Args[1]
	calculator := internal.NewCalculator(internal.NewParser())

	result, err := calculator.Calculate(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %.2f", result)
}

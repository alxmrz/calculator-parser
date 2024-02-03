package internal

import (
	"log"
	"strconv"
)

type Calculator struct {
	parser *Parser
}

func NewCalculator(parser *Parser) *Calculator {
	return &Calculator{
		parser,
	}
}

func (c *Calculator) CalculateOverTree(tree *Node) float64 {
	if tree.left == nil && tree.right != nil {
		val, err := strconv.ParseFloat(tree.right.value, 64)
		if err != nil {
			log.Fatal(err)
		}

		return val
	}
	if tree.left != nil && tree.right == nil {
		val, err := strconv.ParseFloat(tree.left.value, 64)
		if err != nil {
			log.Fatal(err)
		}

		return val
	}

	if tree.left == nil && tree.right == nil {
		val, err := strconv.ParseFloat(tree.value, 64)
		if err != nil {
			log.Fatal(err)
		}

		return val
	}

	return c.calculate(c.CalculateOverTree(tree.left), c.CalculateOverTree(tree.right), tree.value)
}

func (c *Calculator) calculate(lval, rval float64, operation string) float64 {
	switch operation {
	case "+":
		return lval + rval
	case "-":
		return lval - rval
	case "*":
		return lval * rval
	case "/":
		return lval / rval
	default:
		log.Fatal("Unknown operation ", operation)
	}

	return 0
}

func (c *Calculator) Calculate(input string) float64 {
	return c.CalculateOverTree(c.parser.buildTree(input))
}

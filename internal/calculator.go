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

func (c *Calculator) CalculateOverTree(tree *Node) int {
	right, err := strconv.Atoi(tree.right.value)
	if err != nil {
		log.Fatal(err)
	}

	if tree.left == nil {
		return right
	}

	return c.calculate(c.CalculateOverTree(tree.left), right, rune(tree.value[0]))
}

func (c *Calculator) calculate(lval, rval int, operation rune) int {
	switch operation {
	case '+':
		return lval + rval
	case '-':
		return lval - rval
	case '*':
		return lval * rval
	case '/':
		return lval / rval
	default:
		log.Fatal("Unknown operation ", operation)
	}

	return 0
}

func (c *Calculator) Calculate(input string) int {
	return c.CalculateOverTree(c.parser.buildTree(input))
}

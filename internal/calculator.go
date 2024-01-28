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
	if tree.left == nil && tree.right != nil {
		val, err := strconv.Atoi(tree.right.value)
		if err != nil {
			log.Fatal(err)
		}

		return val
	}
	if tree.left != nil && tree.right == nil {
		val, err := strconv.Atoi(tree.left.value)
		if err != nil {
			log.Fatal(err)
		}

		return val
	}

	if tree.left == nil && tree.right == nil {
		val, err := strconv.Atoi(tree.value)
		if err != nil {
			log.Fatal(err)
		}

		return val
	}

	return c.calculate(c.CalculateOverTree(tree.left), c.CalculateOverTree(tree.right), rune(tree.value[0]))
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
		log.Fatal("Unknown operation ", string(operation))
	}

	return 0
}

func (c *Calculator) Calculate(input string) int {
	return c.CalculateOverTree(c.parser.buildTree(input))
}

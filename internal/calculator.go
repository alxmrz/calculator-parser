package internal

import (
	"errors"
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

func (c *Calculator) CalculateOverTree(tree *Node) (float64, error) {
	if tree.left == nil && tree.right != nil {
		val, err := strconv.ParseFloat(tree.right.value, 64)
		if err != nil {
			return 0.0, err
		}

		return val, nil
	}
	if tree.left != nil && tree.right == nil {
		val, err := strconv.ParseFloat(tree.left.value, 64)
		if err != nil {
			return 0.0, err
		}

		return val, nil
	}

	if tree.left == nil && tree.right == nil {
		val, err := strconv.ParseFloat(tree.value, 64)
		if err != nil {
			return 0.0, err
		}

		return val, nil
	}

	left, err := c.CalculateOverTree(tree.left)
	if err != nil {
		return 0.0, err
	}

	right, err := c.CalculateOverTree(tree.right)
	if err != nil {
		return 0.0, err
	}

	return c.calculate(left, right, tree.value)
}

func (c *Calculator) calculate(lval, rval float64, operation string) (float64, error) {
	operations := map[string]func(float64, float64) (float64, error){
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	calcOp, ok := operations[operation]
	if !ok {
		return 0.0, errors.New("Unknown operation " + operation)
	}

	return calcOp(lval, rval)
}

func add(lval, rval float64) (float64, error) {
	return lval + rval, nil
}

func sub(lval, rval float64) (float64, error) {
	return lval - rval, nil
}

func mul(lval, rval float64) (float64, error) {
	return lval * rval, nil
}

func div(lval, rval float64) (float64, error) {
	if rval == 0 {
		return 0.0, errors.New("division by zero")
	}
	return lval / rval, nil
}

func (c *Calculator) Calculate(input string) (float64, error) {
	tree, err := c.parser.buildTree(input)
	if err != nil {
		return 0.0, err
	}

	return c.CalculateOverTree(tree)
}

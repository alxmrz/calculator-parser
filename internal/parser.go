package internal

import (
	"errors"
	"log"
	"strconv"
)

type Parser struct {
}

type Node struct {
	value string
	left  *Node
	right *Node
}

func NewParser() *Parser {
	return &Parser{}
}

// 1 + 1 * 2
func (p *Parser) Parse(input string) (int, error) {
	result := 0
	currentOperation := '+'

	var currentValue string
	for _, token := range input {

		switch token {
		case '+', '-':
			tmp, err := strconv.Atoi(currentValue)
			if err != nil {
				return 0, errors.New("can't atoi")
			}

			result = p.calculate(result, tmp, currentOperation)

			currentValue = ""
			currentOperation = token
			break
		case ' ':
			break
		default:
			currentValue += string(token)
		}

	}

	tmp, err := strconv.Atoi(currentValue)
	if err != nil {
		return 0, errors.New("dsfdf")
	}

	result = p.calculate(result, tmp, currentOperation)

	return result, nil
}

func (p *Parser) calculate(lval, rval int, operation rune) int {
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

func (p *Parser) buildTree(input string) *Node {
	root := &Node{}
	origin := root

	for i := len(input) - 1; i >= 0; i-- {
		tokenString := string(input[i])
		if isOperation(tokenString) {
			root.value = tokenString
			root.left = &Node{}
			root = root.left
			continue
		}

		// skip spaces
		if tokenString == " " {
			continue
		}

		if root.value == "" {
			if root.right == nil {
				root.right = &Node{}
			}

			root.right.value = tokenString + root.right.value
			continue
		}

		if root.left == nil {
			root.left = &Node{}
		}

		root.left.value = tokenString + root.left.value
	}

	return origin
}

func (p *Parser) CalculateOverTree(tree *Node) int {
	right, err := strconv.Atoi(tree.right.value)
	if err != nil {
		log.Fatal(err)
	}

	if tree.left == nil {
		return right
	}

	return p.calculate(p.CalculateOverTree(tree.left), right, rune(tree.value[0]))
}

func isOperation(operation string) bool {
	switch operation {
	case "+", "-", "*", "/":
		return true
	default:
		return false
	}
}

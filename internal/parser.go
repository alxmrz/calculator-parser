package internal

import (
	"errors"
	"log"
	"strconv"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

// (0+1) + 1
func (p *Parser) Parse(input string) (int, error) {
	result := 0
	currentOperation := '+'

	var currentValue string
	for _, token := range input {

		switch token {
		case '+':
			tmp, err := strconv.Atoi(currentValue)
			if err != nil {
				return 0, errors.New("can't atoi")
			}

			result = p.calculate(result, tmp, currentOperation)

			currentValue = ""
			currentOperation = '+'
			break
		case '-':
			tmp, err := strconv.Atoi(currentValue)
			if err != nil {
				return 0, errors.New("can't atoi")
			}

			result = p.calculate(result, tmp, currentOperation)

			currentValue = ""
			currentOperation = '-'
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
	default:
		log.Fatal("Unknown operation ", operation)
	}

	return 0
}

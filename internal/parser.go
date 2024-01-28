package internal

type Parser struct {
}

type Node struct {
	value string
	left  *Node
	right *Node
	prev  *Node
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) buildTree(input string) *Node {
	root := &Node{}
	origin := root

	for i := len(input) - 1; i >= 0; i-- {
		tokenString := string(input[i])
		if isOperation(tokenString) {
			root.value = tokenString
			root.left = &Node{}
			root.prev = root
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

func getWeight(operation string) int {
	switch operation {
	case "+", "-":
		return 10
	case "*", "/":
		return 20
	default:
		return 0
	}
}

func isOperation(operation string) bool {
	switch operation {
	case "+", "-", "*", "/":
		return true
	default:
		return false
	}
}

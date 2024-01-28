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
	//origin := root

	// lval or rval for calculation
	operand := ""

	for i := len(input) - 1; i >= 0; i-- {
		tokenString := string(input[i])

		if isOperation(tokenString) {
			if root.right == nil {
				root.right = &Node{}
			}
			root.right.value = operand
			operand = ""

			root = insertOperationNode(tokenString, root)

			continue
		}

		// skip spaces
		if tokenString == " " {
			continue
		}

		operand = tokenString + operand
	}

	//root.right = &Node{value: operand}
	root.value = operand

	var origin *Node
	for {
		if root.prev == nil {
			origin = root
			break
		}

		root = root.prev
	}

	return origin
}

// 2 + 2 * 3
func insertOperationNode(operation string, root *Node) *Node {
	if root.prev != nil && getWeight(operation) <= getWeight(root.prev.value) {
		return insertOperationNode(operation, root.prev)
	}

	if root.value == "" {
		root.value = operation
		root.left = &Node{
			prev: root,
		}

		return root.left
	}

	if root.prev == nil && getWeight(operation) <= getWeight(root.value) {

		root.prev = &Node{
			right: root,
			value: operation,
		}

		root.prev.left = &Node{
			prev: root.prev,
		}

		return root.prev.left
	}

	return nil
}

func (p *Parser) buildTree1(input string) *Node {
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
	case "+":
		return 10
	case "-":
		return 15
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

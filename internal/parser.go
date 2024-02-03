package internal

type Parser struct {
	multiplier int
}

type Node struct {
	value  string
	left   *Node
	right  *Node
	prev   *Node
	weight int
}

func NewParser() *Parser {
	return &Parser{multiplier: 0}
}

func (p *Parser) buildTree(input string) *Node {
	root := &Node{weight: 0}

	// lval or rval for calculation
	operand := ""

	//Historically I started the parsing from right to left so let it be
	for i := len(input) - 1; i >= 0; i-- {
		tokenString := string(input[i])

		if isOperation(tokenString) {
			if root.right == nil {
				root.right = &Node{weight: 0}
			}
			root.right.value = operand
			operand = ""

			root = p.insertOperationNode(tokenString, root)

			continue
		}

		// skip spaces
		if tokenString == " " {
			continue
		}

		if tokenString == ")" {
			p.multiplier += 100
			continue
		}

		if tokenString == "(" {
			p.multiplier -= 100
			continue
		}

		operand = tokenString + operand
	}

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

func (p *Parser) insertOperationNode(operation string, root *Node) *Node {
	if root.prev != nil && p.getWeight(operation) <= root.prev.weight {
		return p.insertOperationNode(operation, root.prev)
	}

	if root.value == "" {
		root.value = operation
		root.weight = p.getWeight(operation)
		root.left = &Node{
			prev: root,
		}

		return root.left
	}

	if root.prev == nil && p.getWeight(operation) <= root.weight {

		root.prev = &Node{
			right:  root,
			value:  operation,
			weight: p.getWeight(operation),
		}

		root.prev.left = &Node{
			prev: root.prev,
		}

		return root.prev.left
	}

	if p.getWeight(operation) <= root.weight {
	}

	newNode := &Node{
		prev:   root.prev,
		right:  root,
		value:  operation,
		weight: p.getWeight(operation),
		left:   &Node{},
	}

	root.prev.left = newNode

	newNode.left.prev = newNode

	return newNode.left
}

func (p *Parser) getWeight(operation string) int {
	switch operation {
	case "+":
		return 10 + p.multiplier
	case "-":
		return 15 + p.multiplier
	case "*", "/":
		return 20 + p.multiplier
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

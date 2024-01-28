package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildTreeWithPluses(t *testing.T) {
	parser := NewParser()

	expected := &Node{
		value: "+",
		right: &Node{
			value: "+",
			left: &Node{
				right: &Node{
					value: "1",
				},
			},
			right: &Node{
				value: "2",
			},
		},
		left: &Node{
			value: "1",
		},
	}
	actual := parser.buildTree("1 + 1 + 2")

	assert.True(t, AreTreesEqual(expected, actual))
}

func TestBuildTreeWithPlusesAndMultiValues(t *testing.T) {
	parser := NewParser()

	expected := &Node{
		value: "+",
		right: &Node{
			value: "+",
			left: &Node{
				right: &Node{
					value: "15",
				},
			},
			right: &Node{
				value: "20",
			},
		},
		left: &Node{
			value: "10",
		},
	}
	actual := parser.buildTree("10 + 15 + 20")

	assert.True(t, AreTreesEqual(expected, actual))
}

func TestBuildTreeWithMultipliers(t *testing.T) {
	parser := NewParser()

	expected := &Node{
		value: "*",
		left: &Node{
			value: "5",
		},
		right: &Node{
			value: "5",
		},
	}
	actual := parser.buildTree("5 * 5")
	assert.True(t, AreTreesEqual(expected, actual))
}

func TestBuildTreeWithDivisionsMultipliers(t *testing.T) {
	parser := NewParser()

	expected := &Node{
		value: "/",
		left: &Node{
			value: "5",
		},
		right: &Node{
			value: "5",
		},
	}
	actual := parser.buildTree("5 / 5")

	assert.True(t, AreTreesEqual(expected, actual))
}

func TestSortTreeByOrder(t *testing.T) {
	parser := NewParser()

	expected := &Node{
		value: "+",
		right: &Node{
			value: "*",
			left: &Node{
				right: &Node{
					value: "5",
				},
			},
			right: &Node{
				value: "5",
			},
		},
		left: &Node{
			value: "2",
		},
	}

	actual := parser.buildTree("2 + 5 * 5")

	assert.True(t, AreTreesEqual(expected, actual))
}

func AreTreesEqual(t1 *Node, t2 *Node) bool {
	if t1 == nil && t2 != nil || t2 == nil && t1 != nil {
		return false
	}

	if t1 == nil && t2 == nil {
		return true
	}

	if t1.value != t2.value {
		return false
	}

	return AreTreesEqual(t1.left, t2.left) && AreTreesEqual(t1.right, t2.right)
}

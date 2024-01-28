package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		Expected int
		Input    string
	}{
		{2, "1 + 1"},
		{4, "3 + 1"},
		{55, "30 + 25"},
		{0, "1 - 1"},
		{9, "20 - 11"},
		{12, "10 + 5 - 3"},
		{8, "10 - 5 + 3"},
	}

	parser := NewParser()

	for _, test := range tests {
		actual, err := parser.Parse(test.Input)
		assert.NoError(t, err)
		assert.Equal(t, test.Expected, actual)
	}
}

func TestBuildTreeWithPluses(t *testing.T) {
	parser := NewParser()

	expected := &Node{
		value: "+",
		left: &Node{
			value: "+",
			left: &Node{
				right: &Node{
					value: "1",
				},
			},
			right: &Node{
				value: "1",
			},
		},
		right: &Node{
			value: "2",
		},
	}
	actual := parser.buildTree("1 + 1 + 2")

	assert.Equal(t, expected, actual)
}

func TestBuildTreeWithPlusesAndMultiValues(t *testing.T) {
	parser := NewParser()

	expected := &Node{
		value: "+",
		left: &Node{
			value: "+",
			left: &Node{
				right: &Node{
					value: "10",
				},
			},
			right: &Node{
				value: "15",
			},
		},
		right: &Node{
			value: "20",
		},
	}
	actual := parser.buildTree("10 + 15 + 20")

	assert.Equal(t, expected, actual)
}

func TestBuildTreeWithMultipliers(t *testing.T) {
	parser := NewParser()

	expected := &Node{
		value: "*",
		left: &Node{
			right: &Node{
				value: "5",
			},
		},
		right: &Node{
			value: "5",
		},
	}
	actual := parser.buildTree("5 * 5")

	assert.Equal(t, expected, actual)
}

func TestBuildTreeWithDivisionsMultipliers(t *testing.T) {
	parser := NewParser()

	expected := &Node{
		value: "/",
		left: &Node{
			right: &Node{
				value: "5",
			},
		},
		right: &Node{
			value: "5",
		},
	}
	actual := parser.buildTree("5 / 5")

	assert.Equal(t, expected, actual)
}

func _TestSortTreeByOrder(t *testing.T) {
	parser := NewParser()

	expected := &Node{
		value: "+",
		left: &Node{
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
		right: &Node{
			value: "2",
		},
	}

	actual := parser.buildTree("2 + 5 * 5")

	assert.Equal(t, expected, actual)
}

func TestCalculateOverTree(t *testing.T) {
	parser := NewParser()

	tree := &Node{
		value: "+",
		left: &Node{
			value: "+",
			left: &Node{
				right: &Node{
					value: "10",
				},
			},
			right: &Node{
				value: "15",
			},
		},
		right: &Node{
			value: "20",
		},
	}

	expected := 45
	actual := parser.CalculateOverTree(tree)

	assert.Equal(t, expected, actual)
}

func TestCalculateOverTreeDifferentOperations(t *testing.T) {
	parser := NewParser()

	tree := &Node{
		value: "+",
		left: &Node{
			value: "-",
			left: &Node{
				right: &Node{
					value: "10",
				},
			},
			right: &Node{
				value: "15",
			},
		},
		right: &Node{
			value: "20",
		},
	}

	expected := 15
	actual := parser.CalculateOverTree(tree)

	assert.Equal(t, expected, actual)
}

func TestCalculateOverTreeMultipliers(t *testing.T) {
	parser := NewParser()

	tree := &Node{
		value: "*",
		left: &Node{
			right: &Node{
				value: "5",
			},
		},
		right: &Node{
			value: "5",
		},
	}

	expected := 25
	actual := parser.CalculateOverTree(tree)

	assert.Equal(t, expected, actual)
}

func TestCalculateOverTreeDivisions(t *testing.T) {
	parser := NewParser()

	tree := &Node{
		value: "/",
		left: &Node{
			right: &Node{
				value: "5",
			},
		},
		right: &Node{
			value: "5",
		},
	}

	expected := 1
	actual := parser.CalculateOverTree(tree)

	assert.Equal(t, expected, actual)
}

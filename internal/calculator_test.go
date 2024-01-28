package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate(t *testing.T) {
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
		{16, "10 + 2 * 3"},
		{23, "10 * 2 + 3"},
		{22, "2 * 5 + 3 * 4"},
		{25, "5 / 5 + 2 * 3 * 4"},
		{0, "5 * 2 + 40 / 4 - 20"},
		{-5, "0 - 5"},
		{40, "(5+3)*5"},
		{20, "5+3*5"},
		{45, "(2+3)*(4+5)"},
		{95, "5 * ((1+2)*(3+4)) - 10"},
		{17, "50+(1+2*(3+4+(11-8+(15-4)))*(5-6)+7)+1"},
		//{9, "5 + 20 * 2 / 10"}, // TODO: float does not work
	}

	calculator := NewCalculator(NewParser())

	for _, test := range tests {
		actual := calculator.Calculate(test.Input)
		assert.Equal(t, test.Expected, actual)
	}
}

func TestCalculateOverTree(t *testing.T) {
	calculator := NewCalculator(NewParser())

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
	actual := calculator.CalculateOverTree(tree)

	assert.Equal(t, expected, actual)
}

func TestCalculateOverTreeDifferentOperations(t *testing.T) {
	calculator := NewCalculator(NewParser())

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
	actual := calculator.CalculateOverTree(tree)

	assert.Equal(t, expected, actual)
}

func TestCalculateOverTreeMultipliers(t *testing.T) {
	calculator := NewCalculator(NewParser())

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
	actual := calculator.CalculateOverTree(tree)

	assert.Equal(t, expected, actual)
}

func TestCalculateOverTreeDivisions(t *testing.T) {
	calculator := NewCalculator(NewParser())

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
	actual := calculator.CalculateOverTree(tree)

	assert.Equal(t, expected, actual)
}

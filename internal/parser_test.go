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

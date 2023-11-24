package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestURange(t *testing.T) {
	tests := []struct {
		name     string
		min, max int
		expected []int
	}{
		{name: "1", min: 1, max: 1, expected: []int{1}},
		{name: "range of 0, 1, 2", min: 0, max: 2, expected: []int{0, 1, 2}},
		{name: "range of 2, 3, 4, 5", min: 2, max: 5, expected: []int{2, 3, 4, 5}},
		{name: "min maxed swapped when max less than min", min: 5, max: 2, expected: []int{2, 3, 4, 5}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equalf(t, test.expected, Range(test.min, test.max), "Range()")
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "slice of 2", input: []int{1, 2}, expected: []int{2, 1}},
		{name: "slice of 1", input: []int{1}, expected: []int{1}},
		{name: "slice of 0", input: []int{}, expected: []int{}},
		{name: "slice of 9", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equalf(t, test.expected, Reverse(test.input), "Reverse()")
		})
	}
}

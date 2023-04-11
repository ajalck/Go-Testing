package main

import "testing"

func TestCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 0},
		{10, 106},
		{-4, 16},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Errorf("Test Failed :inputted:%v, expected:%v, received:%v", test.input, test.expected, output)
		}
	}
}

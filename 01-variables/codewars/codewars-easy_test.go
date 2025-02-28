package main

import (
	"testing"
)

// MakeUpperCase (8 kyu): Write a function which converts the input string to uppercase
func TestMakeUpperCase(t *testing.T) {

	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "HELLO"},
		{"hello world", "HELLO WORLD"},
		{"hello world !", "HELLO WORLD !"},
		{"heLlO wORLd !", "HELLO WORLD !"},
		{"1,2,3 hello world!", "1,2,3 HELLO WORLD!"},
	}

	for _, test := range tests {
		result := MakeUpperCase(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected '%s' but got '%s'", test.input, test.expected, result)
		}
	}
}

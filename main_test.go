package main

import (
	"reflect"
	"testing"
)

func TestMapFrequency(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{"hello", map[string]int{"h": 1, "e": 1, "l": 2, "o": 1}},
		{"apple", map[string]int{"a": 1, "p": 2, "l": 1, "e": 1}},
		{"aaaa", map[string]int{"a": 4}},
		{"", map[string]int{}},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := MapFrequency(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For input '%s', expected %v but got %v", test.input, test.expected, result)
			}
		})
	}
}

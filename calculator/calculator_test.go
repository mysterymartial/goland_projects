package main

import (
	"reflect"
	"testing"
)

//i want to test

func TestAdd(t *testing.T) {
	result := Add(2, 2)
	expected := 4
	if result != expected {
		t.Errorf("Add(2, 2) = %d; want %d", result, expected)
	}

}
func TestSubtract(t *testing.T) {
	result := subtract(5, 2)
	expected := 3
	if result != expected {
		t.Errorf("Subtract(5,2) = %d; want %d", result, expected)
	}
}
func TestMultiply(t *testing.T) {
	result := multiply(2, 3)
	expected := 6
	if result != expected {
		t.Errorf("Multiply(2,3) = %d; want %d", result, expected)
	}

}
func TestEvenCollection(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3, 4, 5, 6}, expected: []int{2, 4, 6}},
		{input: []int{11, 13, 15}, expected: []int{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := evenNumber(test.input)
			if !reflect.DeepEqual(result, test.expected) {

			}
		})
	}
}
func TestDivision(t *testing.T) {
	result := divide(8, 4)
	expected := 2
	if result != expected {
		t.Errorf("Divide(4,8) = %d; want %d", result, expected)
	}

}
func TestOddCollection(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3, 4, 5, 6}, expected: []int{1, 3, 5}},
		{input: []int{2, 8, 10}, expected: []int{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := oddNumber(test.input)
			if !reflect.DeepEqual(result, test.expected) {

			}
		})
	}

}
func TestReverseArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3, 4, 5, 6}, expected: []int{5, 4, 3, 2, 1}},
		{input: []int{1, 2, 3, 4, 5, 6}, expected: []int{6, 5, 4, 3, 2, 1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := reverseArray(test.input)
			if !reflect.DeepEqual(result, test.expected) {

			}
		})
	}
}
func TestSortedArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{input: []int{5, 6, 3, 1, 4, 2}, expected: []int{1, 2, 3, 4, 5, 6}},
		{input: []int{6, 5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5, 6}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := sortArray(test.input)
			if !reflect.DeepEqual(result, test.expected) {

			}
		})
	}

}
func TestStringPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"palindrome", "bob", true},
		{"palindrome", "cap", false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isPalindrome(test.input)
			if result != test.expected {
			}

		})

	}

}
func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"leapYear", 2023, true},
		{"leapYear", 2023, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isLeapYear(test.input)
			if result != test.expected {

			}
		})
	}

}
func TestFactorial(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"factorial", 5, 720},
		{"factorial", 4, 24},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := factorial(test.input)
			if result != test.expected {

			}
		})
	}

}
func TestFirstDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Duplicate", []int{1, 2, 3, 2, 4, 4}, 2},
		{"Duplicate", []int{1, 2, 3, 4, 5}, -1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := firstDuplicate(test.input)
			if result != test.expected {

			}
		})
	}
}

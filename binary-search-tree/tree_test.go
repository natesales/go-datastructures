package main

import "testing"

var tree = BinarySearchTree{}

func TestInsert(t *testing.T) {
	for i := 0; i < 10; i++ {
		tree.Insert(i)
	}

	for i := -10; i < 0; i++ {
		tree.Insert(i)
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name             string
		value            int
		expected, actual bool
	}{
		{"Root Contains", 0, true, tree.Contains(0)},
		{"Positive-value Contains", 1, true, tree.Contains(1)},
		{"Negative-value Contains", -1, true, tree.Contains(-1)},
		{"Not Contains", 100, false, tree.Contains(100)},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			if test.actual != test.expected {
				t.Errorf("%s: tree contains %d (%t) expected %t", test.name, test.value, test.actual, test.expected)
			}
		})
	}
}

func TestMax(t *testing.T) {
	maxVal := tree.Max()
	if maxVal != 9 {
		t.Errorf("expected Max to be 9, got %d\n", maxVal)
	}
}

func TestMin(t *testing.T) {
	maxVal := tree.Min()
	if maxVal != -10 {
		t.Errorf("expected Min to be -10, got %d\n", maxVal)
	}
}

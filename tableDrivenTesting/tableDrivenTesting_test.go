package tabledriventesting

import "testing"

func add(a, b int) int {
	return a + b
}

func TestTableAdd(t *testing.T) {
	var tests = []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{2, 4, 6},
	}

	for _, test := range tests {
		if res := add(test.a, test.b); res != test.expected {
			t.Error("Test Failed: a:{}, b:{}, res:{}", test.a, test.b, test.expected)
		}
	}
}

package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Inc(x int) (res int) {
	res = x + 1
	return res
}

func TestInc(t *testing.T) {
	if Inc(1) != 2 {
		t.Error("failed")
	}
}

func TestIncWithTestfy(t *testing.T) {
	assert.Equal(t, Inc(1), 2)
}

func TestNotEqWithTestfy(t *testing.T) {
	assert.NotEqual(t, Inc(1), 3)
}

func TestNilEqWithTestfy(t *testing.T) {
	assert.NotNil(t, Inc(1))
}

func TestTableDrivenTest(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    int
		expected int
	}{
		{1, 2},
		{2, 3},
	}

	for _, test := range tests {
		assert.Equal(Inc(test.input), test.expected)
	}
}

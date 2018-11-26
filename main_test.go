// +build unit

package tests

import (
	"fmt"
	"testing"
)

func TestUnit(t *testing.T) {
	fmt.Println("run unit test")
}

// tag:unit
// go test -tags=unit
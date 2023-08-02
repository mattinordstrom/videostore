package main

import "testing"

func TestOneEqualsOne(t *testing.T) {
	var someVar = 1
	var otherVar = 1
	if someVar != otherVar {
		t.Errorf("Expected 1 to equal 1")
	}
}

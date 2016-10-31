package example

import "testing"

// Run this script at the command line "go test -v"

func TestSquare(t *testing.T) {
	if v := Square(4); v != 16 {
		t.Error("expected", 16, "got", v)
	}
}

package p20

import (
	"testing"
)

func TestClosest(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			File:   "a.txt",
			Result: 457,
		},
	}

	for i, test := range tests {
		n := Closest(test.File)
		if n != test.Result {
			t.Fatalf("%d: %s '%d' (should be %d)", i, test.File, n, test.Result)
		}
	}
}

func TestCollide(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			File:   "a.txt",
			Result: 448,
		},
	}

	for i, test := range tests {
		n := Collide(test.File)
		if n != test.Result {
			t.Fatalf("%d: %s '%d' (should be %d)", i, test.File, n, test.Result)
		}
	}
}

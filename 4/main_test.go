package p4

import (
	"testing"
)

func TestValidCount(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			File:   "ta.txt",
			Result: 2,
		},
		{
			File:   "a.txt",
			Result: 451,
		},
	}

	for i, test := range tests {
		n := ValidCount(test.File)
		if n != test.Result {
			t.Fatalf("%d: %d (should be %d)", i, n, test.Result)
		}
	}
}

func TestNewValidCount(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			File:   "tb.txt",
			Result: 3,
		},
		{
			File:   "b.txt",
			Result: 223,
		},
	}

	for i, test := range tests {
		n := NewValidCount(test.File)
		if n != test.Result {
			t.Fatalf("%d: %d (should be %d)", i, n, test.Result)
		}
	}
}

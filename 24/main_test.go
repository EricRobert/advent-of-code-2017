package p24

import (
	"testing"
)

func TestBridge(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			File:   "t.txt",
			Result: 31,
		},
		{
			File:   "a.txt",
			Result: 1511,
		},
	}

	for i, test := range tests {
		n := Bridge(test.File)
		if n != test.Result {
			t.Fatalf("%d: %s '%d' (should be %d)", i, test.File, n, test.Result)
		}
	}
}

func TestLongest(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			File:   "t.txt",
			Result: 19,
		},
		{
			File:   "a.txt",
			Result: 1471,
		},
	}

	for i, test := range tests {
		n := Longest(test.File)
		if n != test.Result {
			t.Fatalf("%d: %s '%d' (should be %d)", i, test.File, n, test.Result)
		}
	}
}

package p18

import (
	"testing"
)

func TestDuet(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			File:   "t1.txt",
			Result: 4,
		},
		{
			File:   "a.txt",
			Result: 1187,
		},
	}

	for i, test := range tests {
		n := Duet(test.File)
		if n != test.Result {
			t.Fatalf("%d: %s '%d' (should be %d)", i, test.File, n, test.Result)
		}
	}
}

func TestTwo(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			File:   "t2.txt",
			Result: 3,
		},
		{
			File:   "a.txt",
			Result: 5969,
		},
	}

	for i, test := range tests {
		n := Two(test.File)
		if n != test.Result {
			t.Fatalf("%d: %s '%d' (should be %d)", i, test.File, n, test.Result)
		}
	}
}

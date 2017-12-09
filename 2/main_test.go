package p2

import "testing"

func TestDiffSum(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			File:   "ta.txt",
			Result: 18,
		},
		{
			File:   "a.txt",
			Result: 43074,
		},
	}

	for i, test := range tests {
		n := DiffSum(test.File)
		if n != test.Result {
			t.Fatalf("%d: %s: %d (should be %d)", i, test.File, n, test.Result)
		}
	}
}

func TestEvenlyDivisibleSum(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			File:   "tb.txt",
			Result: 9,
		},
		{
			File:   "a.txt",
			Result: 280,
		},
	}

	for i, test := range tests {
		n := EvenlyDivisibleSum(test.File)
		if n != test.Result {
			t.Fatalf("%d: %s: %d (should be %d)", i, test.File, n, test.Result)
		}
	}
}

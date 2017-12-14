package p13

import (
	"testing"
)

func TestSeverity(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			File:   "t.txt",
			Result: 24,
		},
		{
			File:   "a.txt",
			Result: 632,
		},
	}

	for i, test := range tests {
		n := Severity(test.File)
		if n != test.Result {
			t.Fatalf("%d: %s '%d' (should be %d)", i, test.File, n, test.Result)
		}
	}
}

func TestDelay(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			File:   "t.txt",
			Result: 10,
		},
		{
			File:   "a.txt",
			Result: 3849742,
		},
	}

	for i, test := range tests {
		n := Delay2(test.File)
		if n != test.Result {
			t.Fatalf("%d: %s '%d' (should be %d)", i, test.File, n, test.Result)
		}
	}
}

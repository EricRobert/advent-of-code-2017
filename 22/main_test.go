package p22

import (
	"testing"
)

func TestInfect(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			File:   "t.txt",
			Result: 5587,
		},
		{
			File:   "a.txt",
			Result: 5280,
		},
	}

	for i, test := range tests {
		n := Infect(test.File)
		if n != test.Result {
			t.Fatalf("%d: %s '%d' (should be %d)", i, test.File, n, test.Result)
		}
	}
}

func TestInfect2(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			File:   "t.txt",
			Result: 2511944,
		},
		{
			File:   "a.txt",
			Result: 2512261,
		},
	}

	for i, test := range tests {
		n := Infect2(test.File)
		if n != test.Result {
			t.Fatalf("%d: %s '%d' (should be %d)", i, test.File, n, test.Result)
		}
	}
}

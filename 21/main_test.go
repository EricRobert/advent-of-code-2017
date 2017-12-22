package p21

import (
	"testing"
)

func TestPixels(t *testing.T) {
	tests := []struct {
		File       string
		Iterations int
		Result     int
	}{
		{
			File:       "t.txt",
			Iterations: 2,
			Result:     12,
		},
		{
			File:       "a.txt",
			Iterations: 5,
			Result:     205,
		},
		{
			File:       "a.txt",
			Iterations: 18,
			Result:     3389823,
		},
	}

	for i, test := range tests {
		n := Pixels(test.File, test.Iterations)
		if n != test.Result {
			t.Fatalf("%d: %s '%d' (should be %d)", i, test.File, n, test.Result)
		}
	}
}

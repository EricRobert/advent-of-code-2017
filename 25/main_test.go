package p25

import (
	"testing"
)

func TestTuring(t *testing.T) {
	tests := []struct {
		Result int
	}{
		{
			Result: 31,
		},
	}

	for i, test := range tests {
		n := Turing()
		if n != test.Result {
			t.Fatalf("%d: '%d' (should be %d)", i, n, test.Result)
		}
	}
}

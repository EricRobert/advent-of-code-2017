package p5

import (
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		Table  []int
		Result int
	}{
		{
			Table:  []int{0, 3, 0, 1, -3},
			Result: 5,
		},
	}

	for i, test := range tests {
		n := Run(test.Table)
		if n != test.Result {
			t.Fatalf("%d: %d (should be %d)", i, n, test.Result)
		}
	}
}

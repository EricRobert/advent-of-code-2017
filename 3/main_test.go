package p3

import "testing"

func TestSteps(t *testing.T) {
	tests := []struct {
		N      int
		Result int
	}{
		{
			N:      1,
			Result: 0,
		},
		{
			N:      12,
			Result: 3,
		},
		{
			N:      23,
			Result: 2,
		},
		{
			N:      1024,
			Result: 31,
		},
		{
			N:      347991,
			Result: 480,
		},
	}

	for i, test := range tests {
		n := Steps(test.N)
		if n != test.Result {
			t.Fatalf("%d: %d: %d (should be %d)", i, test.N, n, test.Result)
		}
	}
}

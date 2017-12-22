package p19

import (
	"testing"
)

func TestWalk(t *testing.T) {
	tests := []struct {
		File   string
		Result string
		Steps  int
	}{
		{
			File:   "t.txt",
			Result: "ABCDEF",
			Steps:  38,
		},
		{
			File:   "a.txt",
			Result: "VTWBPYAQFU",
			Steps:  17358,
		},
	}

	for i, test := range tests {
		s, n := Walk(test.File)
		if s != test.Result || n != test.Steps {
			t.Fatalf("%d: %s '%s' %d (should be %s %d)", i, test.File, s, n, test.Result, test.Steps)
		}
	}
}

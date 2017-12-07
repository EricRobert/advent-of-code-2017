package p7

import (
	"testing"
)

func TestRoot(t *testing.T) {
	tests := []struct {
		File   string
		Result string
	}{
		{
			File:   "t.txt",
			Result: "tknk",
		},
		{
			File:   "a.txt",
			Result: "coco",
		},
	}

	for i, test := range tests {
		s := Root(test.File)
		if s != test.Result {
			t.Fatalf("%d: '%s' (should be %s)", i, s, test.Result)
		}
	}
}

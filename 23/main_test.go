package p23

import (
	"testing"
)

func TestMuls(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			File:   "a.txt",
			Result: 3025,
		},
	}

	for i, test := range tests {
		n := Muls(test.File)
		if n != test.Result {
			t.Fatalf("%d: %s '%d' (should be %d)", i, test.File, n, test.Result)
		}
	}
}

func TestDebug(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			Result: 915,
		},
	}

	for i, test := range tests {
		n := Debug()
		if n != test.Result {
			t.Fatalf("%d: '%d' (should be %d)", i, n, test.Result)
		}
	}
}

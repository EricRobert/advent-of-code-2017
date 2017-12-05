package p4

import (
	"strings"
	"testing"
)

func TestValidCount(t *testing.T) {
	tests := []struct {
		Lines  string
		Result int
	}{
		{
			Lines: `
aa bb cc dd ee
aa bb cc dd aa
aa bb cc dd aaa`,
			Result: 2,
		},
	}

	for i, test := range tests {
		n := ValidCount(strings.Split(test.Lines, "\n"))
		if n != test.Result {
			t.Fatalf("%d: %d (should be %d)", i, n, test.Result)
		}
	}
}

func TestNewValidCount(t *testing.T) {
	tests := []struct {
		Lines  string
		Result int
	}{
		{
			Lines: `
abcde fghij
abcde xyz ecdab
a ab abc abd abf abj
iiii oiii ooii oooi oooo
oiii ioii iioi iiio`,
			Result: 3,
		},
	}

	for i, test := range tests {
		n := NewValidCount(strings.Split(test.Lines, "\n"))
		if n != test.Result {
			t.Fatalf("%d: %d (should be %d)", i, n, test.Result)
		}
	}
}

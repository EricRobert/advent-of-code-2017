package p16

import (
	"testing"
)

func TestDance(t *testing.T) {
	tests := []struct {
		File   string
		Key    string
		Result string
	}{
		{
			File:   "t.txt",
			Key:    "abcde",
			Result: "baedc",
		},
		{
			File:   "a.txt",
			Key:    "abcdefghijklmnop",
			Result: "cknmidebghlajpfo",
		},
	}

	for i, test := range tests {
		s := Dance(test.File, test.Key)
		if s != test.Result {
			t.Fatalf("%d: %s '%s' (should be %s)", i, test.File, s, test.Result)
		}
	}
}

func TestLotsOfDances(t *testing.T) {
	tests := []struct {
		File   string
		Key    string
		Result string
	}{
		{
			File:   "a.txt",
			Key:    "abcdefghijklmnop",
			Result: "cbolhmkgfpenidaj",
		},
	}

	for i, test := range tests {
		s := LotsOfDances(test.File, test.Key)
		if s != test.Result {
			t.Fatalf("%d: %s '%s' (should be %s)", i, test.File, s, test.Result)
		}
	}
}

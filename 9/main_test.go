package p9

import (
	"testing"
)

func TestScore(t *testing.T) {
	tests := []struct {
		Stream string
		Result int
	}{
		{
			Stream: "{}",
			Result: 1,
		},
		{
			Stream: "{{{}}}",
			Result: 6,
		},
		{
			Stream: "{{},{}}",
			Result: 5,
		},
		{
			Stream: "{{{},{},{{}}}}",
			Result: 16,
		},
		{
			Stream: "{<a>,<a>,<a>,<a>}",
			Result: 1,
		},
		{
			Stream: "{{<ab>},{<ab>},{<ab>},{<ab>}}",
			Result: 9,
		},
		{
			Stream: "{{<!!>},{<!!>},{<!!>},{<!!>}}",
			Result: 9,
		},
		{
			Stream: "{{<a!>},{<a!>},{<a!>},{<ab>}}",
			Result: 3,
		},
		{
			Stream: "{{{<\"\"!>,<!<'}'ui!!!>!!!>!<{!!!!!>>},{}}}",
			Result: 9,
		},
	}

	for i, test := range tests {
		n := Score(test.Stream)
		if n != test.Result {
			t.Fatalf("%d: %s '%d' (should be %d)", i, test.Stream, n, test.Result)
		}
	}
}

func TestGarbage(t *testing.T) {
	tests := []struct {
		Stream string
		Result int
	}{
		{
			Stream: "{}",
			Result: 0,
		},
		{
			Stream: "<random characters>",
			Result: 17,
		},
		{
			Stream: "<<<<>",
			Result: 3,
		},
		{
			Stream: "<{!>}>",
			Result: 2,
		},
		{
			Stream: "<!!>",
			Result: 0,
		},
		{
			Stream: "<!!!>>",
			Result: 0,
		},
		{
			Stream: "<{o\"i!a,<{i<a>",
			Result: 10,
		},
	}

	for i, test := range tests {
		n := Garbage(test.Stream)
		if n != test.Result {
			t.Fatalf("%d: %s '%d' (should be %d)", i, test.Stream, n, test.Result)
		}
	}
}

package p1

import (
	"fmt"
	"log"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 1[a|b] 'list-of-digits'")
	}

	digits := args[1]

	n := 0
	switch args[0] {
	case "1a", "1":
		n = Sum(digits, 1)
	case "1b":
		// make it work with any utf-8 string
		half := 0
		for range digits {
			half++
		}

		n = Sum(digits, half/2)
	}

	fmt.Printf("%s: %d\n", digits, n)
}

func Sum(s string, steps int) int {
	get := func(i int) (result rune, next int) {
		// let's not assume that the string is ASCII; should work with any utf-8 string
		for k, c := range s[i:] {
			if k != 0 {
				next = i + k
				break
			}

			result = c
		}

		return
	}

	match := func(i int) bool {
		a, j := get(i)
		b, _ := get((j + steps - 1) % len(s))
		return a == b
	}

	sum := 0
	for i, c := range s {
		if match(i) {
			sum += int(c) - int('0')
		}
	}

	return sum
}

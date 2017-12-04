package main

import (
	"log"
	"os"

	p1 "./1"
	p2 "./2"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("usage: advent-of-code-2017 day [args...]")
	}

	day, args := args[1], args[1:]

	switch day {
	case "1a", "1":
		p1.Main(args)
	case "1b":
		p1.Main(args)
	case "2":
		p2.Main(args)
	default:
		log.Fatalf("%s is unknown", day)
	}
}

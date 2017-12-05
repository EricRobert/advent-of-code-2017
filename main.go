package main

import (
	"log"
	"os"

	p1 "./1"
	p2 "./2"
	p3 "./3"
	p4 "./4"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("usage: advent-of-code-2017 day [args...]")
	}

	day, args := args[1], args[1:]

	switch day {
	case "1a", "1b", "1":
		p1.Main(args)
	case "2a", "2b", "2":
		p2.Main(args)
	case "3a", "3b", "3":
		p3.Main(args)
	case "4a", "4b", "4":
		p4.Main(args)
	default:
		log.Fatalf("%s is unknown", day)
	}
}

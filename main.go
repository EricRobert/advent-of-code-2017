package main

import (
	"log"
	"os"

	p1 "./1"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("usage: advent-of-code-2017 day [args...]")
	}

	day, args := args[1], args[2:]

	switch day {
	case "1":
		p1.Main(args)
	default:
		log.Fatalf("%s is unknown", day)
	}
}

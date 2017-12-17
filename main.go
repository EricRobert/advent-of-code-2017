package main

import (
	"log"
	"os"

	p1 "./1"
	p10 "./10"
	p11 "./11"
	p12 "./12"
	p13 "./13"
	p14 "./14"
	p15 "./15"
	p16 "./16"
	p2 "./2"
	p3 "./3"
	p4 "./4"
	p5 "./5"
	p6 "./6"
	p7 "./7"
	p8 "./8"
	p9 "./9"
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
	case "5a", "5b", "5":
		p5.Main(args)
	case "6a", "6b", "6":
		p6.Main(args)
	case "7a", "7b", "7":
		p7.Main(args)
	case "8a", "8b", "8":
		p8.Main(args)
	case "9a", "9b", "9":
		p9.Main(args)
	case "10a", "10b", "10":
		p10.Main(args)
	case "11a", "11b", "11":
		p11.Main(args)
	case "12a", "12b", "12":
		p12.Main(args)
	case "13a", "13b", "13":
		p13.Main(args)
	case "14a", "14b", "14":
		p14.Main(args)
	case "15a", "15b", "15":
		p15.Main(args)
	case "16a", "16b", "16":
		p16.Main(args)
	default:
		log.Fatalf("%s is unknown", day)
	}
}

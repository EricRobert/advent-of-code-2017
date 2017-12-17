package p16

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2017 16[a|b] filename key")
	}

	switch args[0] {
	case "16a", "16":
		fmt.Print(Dance(args[1], args[2]))
	case "16b":
		fmt.Print(LotsOfDances(args[1], args[2]))
	}
}

func play(filename string, p []int) []int {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	q := make([]int, len(p))

	for _, m := range strings.Split(string(f), ",") {
		if m == "" {
			continue
		}

		switch m[0] {
		case 's':
			n, err := strconv.Atoi(m[1:])
			if err != nil {
				log.Fatal(err)
			}

			for i := 0; i < len(p); i++ {
				j := (i + n) % len(p)
				q[j] = p[i]
			}

			p, q = q, p
		case 'x':
			num := func(s string) int {
				i, err := strconv.Atoi(s)
				if err != nil {
					log.Fatal(err)
				}
				return i
			}

			i := strings.Index(m, "/")
			a := num(m[1:i])
			b := num(m[i+1:])
			p[a], p[b] = p[b], p[a]
		case 'p':
			find := func(c byte) int {
				k := int(c)
				for i := range p {
					if p[i] == k {
						return i
					}
				}

				return -1
			}

			a := find(m[1])
			b := find(m[3])
			p[a], p[b] = p[b], p[a]
		}

		//fmt.Println(m, pretty(p))
	}

	return p
}

func pretty(x []int) string {
	r := make([]byte, len(x))
	for i := range x {
		r[i] = byte(x[i])
	}

	return string(r)
}

func Dance(filename, key string) string {
	p := make([]int, 0)

	for _, c := range key {
		p = append(p, int(c))
	}

	p = play(filename, p)
	return pretty(p)
}

func LotsOfDances(filename, key string) string {
	p := make([]int, 0)

	for _, c := range key {
		p = append(p, int(c))
	}

	m := map[string]int{
		key: 0,
	}

	for {
		p = play(filename, p)

		r := pretty(p)
		if _, ok := m[r]; ok {
			break
		}

		m[r] = len(m)
	}

	n := len(m)
	log.Println("loop:", pretty(p), n)

	for i := 0; i < 1000000000%n; i++ {
		p = play(filename, p)
	}

	return pretty(p)
}

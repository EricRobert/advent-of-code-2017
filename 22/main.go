package p22

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2017 22[a|b] filename")
	}

	switch args[0] {
	case "22a", "22":
		fmt.Print(Infect(args[1]))
	}
}

func load(filename string) (map[string]struct{}, int, int) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[string]struct{})

	y := 0
	for _, line := range strings.Split(string(f), "\n") {
		if line == "" {
			continue
		}

		for x := range line {
			if line[x] == '#' {
				m[fmt.Sprintf("%dx%d", x, y)] = struct{}{}
			}
		}

		y++
	}

	return m, y / 2, y / 2
}

func Infect(filename string) int {
	m, x, y := load(filename)

	key := func(x, y int) string {
		return fmt.Sprintf("%dx%d", x, y)
	}

	count := 0

	dir := 0
	for burst := 0; burst < 10000; burst++ {
		if _, ok := m[key(x, y)]; ok {
			dir++
			delete(m, key(x, y))
		} else {
			if dir--; dir < 0 {
				dir += 4
			}

			m[key(x, y)] = struct{}{}
			count++
		}

		switch dir % 4 {
		case 0:
			y--
		case 1:
			x++
		case 2:
			y++
		case 3:
			x--
		}
	}

	return count
}

const (
	Weakened = 1
	Infected = 2
	Flagged  = 3
)

func Infect2(filename string) int {
	m0, x, y := load(filename)

	m := make(map[string]int)
	for k := range m0 {
		m[k] = Infected
	}

	key := func(x, y int) string {
		return fmt.Sprintf("%dx%d", x, y)
	}

	count := 0

	dir := 0
	for burst := 0; burst < 10000000; burst++ {
		state, ok := m[key(x, y)]
		if !ok {
			if dir--; dir < 0 {
				dir += 4
			}

			m[key(x, y)] = Weakened
		} else {
			switch state {
			case Weakened:
				m[key(x, y)] = Infected
				count++
			case Infected:
				dir++
				m[key(x, y)] = Flagged
			case Flagged:
				dir += 2
				delete(m, key(x, y))
			}
		}

		switch dir % 4 {
		case 0:
			y--
		case 1:
			x++
		case 2:
			y++
		case 3:
			x--
		}
	}

	return count
}

package p20

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2017 20[a|b] filename")
	}

	switch args[0] {
	case "20a", "20":
		fmt.Print(Closest(args[1]))
	case "20b":
		fmt.Print(Collide(args[1]))
	}
}

type Particle struct {
	P []int
	V []int
	A []int
}

func (p *Particle) Step() {
	p.V[0] += p.A[0]
	p.V[1] += p.A[1]
	p.V[2] += p.A[2]
	p.P[0] += p.V[0]
	p.P[1] += p.V[1]
	p.P[2] += p.V[2]
}

func (p *Particle) DistanceAt(t int) int {
	d := func(p, v, a int) int {
		return p + v*t + a*t*t/2
	}

	dx := d(p.P[0], p.V[0], p.A[0])
	dy := d(p.P[1], p.V[1], p.A[1])
	dz := d(p.P[2], p.V[2], p.A[2])

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	return abs(dx) + abs(dy) + abs(dz)
}

func load(filename string) []*Particle {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	m := make([]*Particle, 0)

	for _, line := range strings.Split(string(f), "\n") {
		if line == "" {
			continue
		}

		parse := func(k int) ([]int, int) {
			i := k + strings.Index(line[k:], "<")
			j := i + strings.Index(line[i+1:], ">")
			data := line[i+1 : j+1]
			w := make([]int, 0)
			for _, n := range strings.Split(data, ",") {
				f, err := strconv.Atoi(n)
				if err != nil {
					log.Fatal(err)
				}
				w = append(w, f)
			}

			return w, j + 1
		}

		p, i := parse(0)
		v, i := parse(i)
		a, i := parse(i)

		m = append(m, &Particle{P: p, V: v, A: a})
	}

	return m
}

func Closest(filename string) int {
	m := load(filename)

	closest, minimum := -1, 0
	for i, p := range m {
		d := p.DistanceAt(1000)
		if closest == -1 || d < minimum {
			minimum = d
			closest = i
		}
	}

	return closest
}

func Collide(filename string) int {
	m := load(filename)

	for n := 0; n < 1000; n++ {
		z := make(map[string][]int)

		for i, p := range m {
			if p == nil {
				continue
			}

			p.Step()
			k := fmt.Sprintf("%d,%d,%d", p.P[0], p.P[1], p.P[2])
			z[k] = append(z[k], i)
		}

		for _, list := range z {
			if len(list) == 1 {
				continue
			}

			for _, i := range list {
				m[i] = nil
			}
		}
	}

	count := 0
	for _, p := range m {
		if p != nil {
			count++
		}
	}

	return count
}

package p18

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"sync"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2017 18[a|b] filename")
	}

	switch args[0] {
	case "18a", "18":
		fmt.Print(Duet(args[1]))
	case "18b":
		fmt.Print(Two(args[1]))
	}
}

func load(filename string) [][]string {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	ops := make([][]string, 0)

	for _, line := range strings.Split(string(f), "\n") {
		if line == "" {
			continue
		}

		ops = append(ops, strings.Fields(line))
	}

	return ops
}

func Duet(filename string) int {
	ops := load(filename)
	ip, regs, f := 0, make([]int, 26), 0

	value := func(s string) int {
		if n := s[0]; n >= 'a' && n <= 'z' {
			return regs[n-'a']
		}

		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		return i
	}

	set := func(s string, i int) {
		regs[s[0]-'a'] = i
	}

	for ip >= 0 && ip < len(ops) {
		op := ops[ip]
		ip++

		switch op[0] {
		case "snd":
			f = value(op[1])
		case "set":
			set(op[1], value(op[2]))
		case "add":
			set(op[1], value(op[1])+value(op[2]))
		case "mul":
			set(op[1], value(op[1])*value(op[2]))
		case "mod":
			set(op[1], value(op[1])%value(op[2]))
		case "rcv":
			if op[1] != "0" {
				return f
			}
		case "jgz":
			if value(op[1]) > 0 {
				ip += value(op[2]) - 1
			}
		default:
			log.Fatal("unknown op: " + strings.Join(op, " "))
		}
	}

	return 0
}

type watch struct {
	c chan func()
	n []int
}

func (w *watch) deadlock() bool {
	for _, k := range w.n {
		if k != -1 {
			return false
		}
	}

	return true
}

type queues struct {
	w *watch
	q []chan int
}

func (q *queues) get(p int) (int, bool) {
	c := make(chan bool)
	q.w.c <- func() {
		q.w.n[p]--
		c <- q.w.deadlock()
	}

	if d := <-c; d {
		close(q.q[(p+1)%len(q.q)])
		return 0, false
	}

	i := <-q.q[p]
	return i, true
}

func (q *queues) set(p, i int) {
	j := (p + 1) % len(q.q)
	q.q[j] <- i
	q.w.c <- func() {
		q.w.n[j]++
	}
}

func Two(filename string) int {
	ops := load(filename)

	q := &queues{
		w: &watch{
			c: make(chan func()),
			n: make([]int, 2),
		},
		q: []chan int{
			make(chan int, 65536),
			make(chan int, 65536),
		},
	}

	go func() {
		for f := range q.w.c {
			f()
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(2)
	defer wg.Wait()

	c := make(chan int)

	go func() {
		process(0, ops, q)
		wg.Done()
	}()

	go func() {
		c <- process(1, ops, q)
		wg.Done()
	}()

	return <-c
}

func process(p int, ops [][]string, q *queues) int {
	ip, regs, n := 0, make([]int, 26), 0

	regs['p'-'a'] = p

	value := func(s string) int {
		if n := s[0]; n >= 'a' && n <= 'z' {
			return regs[n-'a']
		}

		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		return i
	}

	set := func(s string, i int) {
		regs[s[0]-'a'] = i
	}

	for ip >= 0 && ip < len(ops) {
		op := ops[ip]
		ip++

		switch op[0] {
		case "snd":
			q.set(p, value(op[1]))
			n++
		case "set":
			set(op[1], value(op[2]))
		case "add":
			set(op[1], value(op[1])+value(op[2]))
		case "mul":
			set(op[1], value(op[1])*value(op[2]))
		case "mod":
			set(op[1], value(op[1])%value(op[2]))
		case "rcv":
			k, ok := q.get(p)
			if !ok {
				return n
			}

			set(op[1], k)
		case "jgz":
			if value(op[1]) > 0 {
				ip += value(op[2]) - 1
			}
		default:
			log.Fatal("unknown op: " + strings.Join(op, " "))
		}
	}

	return n
}

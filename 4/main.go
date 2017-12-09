package p4

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 4[a|b] 'filename'")
	}

	switch args[0] {
	case "4a", "4":
		fmt.Println(ValidCount(args[1]))
	case "4b":
		fmt.Println(NewValidCount(args[1]))
	}
}

func load(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("%s: %s", filename, err)
	}

	return strings.Split(string(data), "\n")
}

func Valid(text string) bool {
	m := make(map[string]struct{})

	for _, p := range strings.Fields(text) {
		if _, ok := m[p]; ok {
			return false
		}

		m[p] = struct{}{}
	}

	return true
}

func ValidCount(filename string) int {
	list := load(filename)

	total := 0

	for _, item := range list {
		if item == "" {
			continue
		}

		if Valid(item) {
			total++
		}
	}

	return total
}

func NewValid(text string) bool {
	m := make(map[string]struct{})

	for _, p := range strings.Fields(text) {
		r := []rune(p)

		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})

		s := string(r)

		if _, ok := m[s]; ok {
			return false
		}

		m[s] = struct{}{}
	}

	return true
}

func NewValidCount(filename string) int {
	list := load(filename)

	total := 0

	for _, item := range list {
		if item == "" {
			continue
		}

		if NewValid(item) {
			total++
		}
	}

	return total
}

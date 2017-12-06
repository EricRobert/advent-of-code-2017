package p6

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 5 'filename'")
	}

	ints := make([]int, 0)

	for _, k := range strings.Fields(args[1]) {
		i, err := strconv.Atoi(k)
		if err != nil {
			log.Fatal(err)
		}

		ints = append(ints, i)
	}

	fmt.Print(Run(ints))
}

func Run(banks []int) int {
	m := make(map[string]struct{})

	hit := func() bool {
		i := fmt.Sprintf("%v", banks)
		if _, ok := m[i]; ok {
			return true
		}

		m[i] = struct{}{}
		return false
	}

	max := func() (int, int) {
		k, n := 0, banks[0]

		for i, p := range banks {
			if p > n {
				k, n = i, p
			}
		}

		return k, n
	}

	for {
		k, n := max()

		banks[k] = 0

		for i := 0; i != n; i++ {
			j := k + i + 1
			banks[j%len(banks)]++
		}

		if hit() {
			break
		}
	}

	return len(m) + 1
}

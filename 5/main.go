package p5

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 5 'filename'")
	}

	data, err := ioutil.ReadFile(args[1])
	if err != nil {
		log.Fatalf("%s: %s", args[1], err)
	}

	ints := make([]int, 0)

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		ints = append(ints, i)
	}

	fmt.Print(Run(ints))
}

func Run(jumps []int) int {
	n, i := 0, 0

	for {
		k := jumps[i]
		jumps[i]++
		n++
		i += k

		if i < 0 || i >= len(jumps) {
			break
		}
	}

	return n
}

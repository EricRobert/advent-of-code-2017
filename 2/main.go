package p2

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 2 'file.tsv'")
	}

	data, err := ioutil.ReadFile(args[1])
	if err != nil {
		log.Fatalf("%s: %s", args[1], err)
	}

	n := Checksum(Parse(string(data)))
	fmt.Println(n)
}

func Parse(data string) [][]int {
	r := make([][]int, 0)

	s := bufio.NewScanner(strings.NewReader(data))
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if len(line) == 0 {
			continue
		}

		ints := make([]int, 0)
		for i, f := range strings.Split(line, "\t") {
			n, err := strconv.Atoi(f)
			if err != nil {
				log.Fatalf("line %d: token %s: %s", i, f, err)
			}

			ints = append(ints, n)
		}

		r = append(r, ints)
	}

	return r
}

func Checksum(rows [][]int) int {
	diff := func(row []int) int {
		if len(row) == 0 {
			return 0
		}

		min, max := row[0], row[0]

		for _, k := range row {
			if k > max {
				max = k
			}

			if k < min {
				min = k
			}
		}

		return max - min
	}

	sum := 0

	for _, row := range rows {
		sum += diff(row)
	}

	return sum
}

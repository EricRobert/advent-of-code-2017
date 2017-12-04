package p2

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 2[a|b] 'file.tsv'")
	}

	data, err := ioutil.ReadFile(args[1])
	if err != nil {
		log.Fatalf("%s: %s", args[1], err)
	}

	n, p := 0, Parse(string(data))

	switch args[0] {
	case "2a", "2":
		n = DiffSum(p)
	case "2b":
		n = EvenlyDivisibleSum(p)
	}

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

func DiffSum(rows [][]int) int {
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

func EvenlyDivisibleSum(rows [][]int) int {
	check := func(row []int) int {
		sort.Ints(row)

		total := 0

		for i, x := range row {
			for j := 0; j != i; j++ {
				y := row[j]
				if k := x / y; k*y == x {
					total += k
				}

				if y*2 >= x {
					break
				}
			}
		}

		return total
	}

	sum := 0

	for _, row := range rows {
		sum += check(row)
	}

	return sum
}

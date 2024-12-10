package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	ls, rs := parseInput(input)

	if part2 {
		r_counts := map[int]int{}

		for _, b := range rs {
			r_counts[b]++
		}

		similarity := 0

		for _, a := range ls {
			similarity += a * r_counts[a]
		}

		return similarity
	} else {
		sort.Ints(ls)
		sort.Ints(rs)

		distance := 0

		for i, l := range ls {
			r := rs[i]

			distance += abs(r - l)
		}

		return distance
	}

}

func parseInput(input string) ([]int, []int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	ls, rs := []int{}, []int{}

	for _, line := range lines {
		l, r := parseLn(line)

		ls = append(ls, l)
		rs = append(rs, r)
	}
	return ls, rs
}

func parseLn(line string) (int, int) {
	pair := strings.Split(line, "   ")

	var l, r int

	if len(pair) != 2 {
		log.Panicf("invalid input %q", line)
	}

	_, err := fmt.Sscanf(line, "%d    %d", &l, &r)

	if err != nil {
		log.Panicf("invalid input %q", line)
	}

	return l, r
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

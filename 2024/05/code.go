package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type orderingRule struct {
	l, r int
}

type update struct {
	middle int
	lookup map[int]int
}

func (u update) follows(rules []orderingRule) bool {
	for _, rule := range rules {
		l, ok := u.lookup[rule.l]

		if !ok {
			continue
		}

		r, ok := u.lookup[rule.r]

		if !ok {
			continue
		}

		if l > r {
			return false
		}
	}

	return true
}

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}

	sections := strings.Split(strings.TrimSpace(input), "\n\n")

	if len(sections) != 2 {
		panic("unexpected input format")
	}

	orderingRules := []orderingRule{}

	for _, line := range strings.Split(strings.TrimSpace(sections[0]), "\n") {
		rawEdge := strings.Split(line, "|")

		orderingRules = append(orderingRules, orderingRule{parseInt(rawEdge[0]), parseInt(rawEdge[1])})
	}

	updates := []update{}

	for _, line := range strings.Split(strings.TrimSpace(sections[1]), "\n") {
		raw := []int{}
		lookup := map[int]int{}

		for i, page := range strings.Split(line, ",") {
			n := parseInt(page)
			raw = append(raw, n)
			lookup[n] = i
		}

		updates = append(updates, update{middleValue(raw), lookup})
	}

	sumOfMiddles := 0
	for _, update := range updates {
		if update.follows(orderingRules) {
			sumOfMiddles += update.middle
		}
	}

	// solve part 1 here
	return sumOfMiddles
}

func middleValue(raw []int) int {
	i := len(raw) / 2
	return raw[i]
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

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
	raw    []int
	lookup map[int]int
}

// part 1
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

// part 2
func (u update) correctedMiddle(rules []orderingRule) int {
	filteredRules := []orderingRule{}

	for _, rule := range rules {
		_, ok := u.lookup[rule.l]
		if !ok {
			continue
		}

		_, ok = u.lookup[rule.r]
		if !ok {
			continue
		}
		filteredRules = append(filteredRules, rule)
	}

	sorted := topologicalSort(u.raw, filteredRules)

	return middleValue(sorted)
}

func topologicalSort(updates []int, filteredRules []orderingRule) []int {
	// Setup
	adjList := make(map[int][]int, len(updates))
	visited := make(map[int]bool, len(updates))

	for _, update := range updates {
		adjList[update] = []int{}
		visited[update] = false
	}

	for _, rule := range filteredRules {
		adjList[rule.l] = append(adjList[rule.l], rule.r)
	}

	// Run topological sort
	stack := []int{}

	for _, update := range updates {
		if !visited[update] {
			topologicalSortHelper(update, adjList, &visited, &stack)
		}
	}

	return stack
}

func topologicalSortHelper(update int, adjList map[int][]int, visited *map[int]bool, stack *[]int) {
	(*visited)[update] = true

	for _, neighbor := range adjList[update] {
		if !(*visited)[neighbor] {
			topologicalSortHelper(neighbor, adjList, visited, stack)
		}
	}

	*stack = append(*stack, update)
}

func run(part2 bool, input string) any {
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

		updates = append(updates, update{raw, lookup})
	}

	sumOfMiddles := 0
	for _, update := range updates {
		if part2 {
			if !update.follows(orderingRules) {
				sumOfMiddles += update.correctedMiddle(orderingRules)
			}

		} else {
			if update.follows(orderingRules) {
				sumOfMiddles += middleValue(update.raw)
			}
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

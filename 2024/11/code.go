package main

import (
	"maps"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	stones := map[int]int{}

	for _, v := range strings.Split(strings.TrimSpace(input), " ") {
		stones[parseInt(v)]++
	}

	blinks := 25
	if part2 {
		blinks = 75
	}

	total := 0
	for count := range maps.Values(simulateBlink(stones, blinks)) {
		total += count
	}
	return total

}

func simulateBlink(stones map[int]int, remainingBlinks int) map[int]int {
	// key value store where the key is the number on the stone and value is count of occurences
	// Since we're getting tons of repeat #s we can dedupe the amount of calculates we do
	nextStones := map[int]int{}

	for k, count := range maps.All(stones) {
		if k == 0 {
			nextStones[1] += count
		} else if isEven, l, r := splitDigits(k); isEven {
			nextStones[l] += count
			nextStones[r] += count
		} else {
			nextStones[k*2024] += count
		}
	}
	if remainingBlinks == 1 {
		return nextStones
	} else {
		return simulateBlink(nextStones, remainingBlinks-1)
	}
}

func splitDigits(value int) (bool, int, int) {
	if value < 10 {
		return false, -1, -1
	}
	stringValue := strconv.Itoa(value)
	isEven := len(stringValue)%2 == 0
	midPoint := len(stringValue) / 2
	return isEven, parseInt(stringValue[:midPoint]), parseInt(stringValue[midPoint:])

}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

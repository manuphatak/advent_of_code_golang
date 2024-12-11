package main

import (
	"regexp"
	"strconv"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}

	re := regexp.MustCompile(`(?m)(?P<mul>mul\((?P<l>\d{0,3}),(?P<r>\d{0,3})\))`)

	lIndex, rIndex := re.SubexpIndex("l"), re.SubexpIndex("r")

	sumProduct := 0
	for _, match := range re.FindAllStringSubmatch(input, -1) {
		l, r := parseInt(match[lIndex]), parseInt(match[rIndex])
		sumProduct += l * r
	}

	return sumProduct
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

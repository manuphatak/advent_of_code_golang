package main

import (
	"strconv"
	"strings"

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
	lines := strings.Split(strings.TrimSpace(input), "\n")
	reports := [][]int{}
	for _, line := range lines {
		report := parseReport(line)

		reports = append(reports, report)
	}

	count := 0
	for _, report := range reports {
		if part2 {
			if isSafePart2(report) {
				count++
			}
		} else {
			if isSafePart1(report) {
				count++
			}
		}
	}

	return count
}

func parseReport(line string) []int {
	report := []int{}

	for _, num := range strings.Split(line, " ") {
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		report = append(report, n)
	}
	return report
}

func isSafePart1(report []int) bool {
	direction := 0

	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]

		switch direction {
		case 0:
			{
				if diff > 0 {
					direction = 1
				}
				if diff < 0 {
					direction = -1
				}
			}
		case 1:
			if diff < 0 {
				return false
			}
		case -1:
			if diff > 0 {
				return false
			}

		}

		absDiff := abs(diff)
		if absDiff <= 0 || absDiff > 3 {
			return false
		}
	}
	return true
}

func isSafePart2(report []int) bool {
	for dampened := -1; dampened < len(report); dampened++ {
		if safeWithDampener(report, dampened) {
			return true
		}
	}

	return false
}

func safeWithDampener(report []int, dampened int) bool {
	direction := 0

	for i := 0; i < len(report)-1; i++ {

		if dampened == 0 && i == 0 {
			continue
		}
		if dampened == len(report)-1 && i == len(report)-2 {
			continue
		}
		a, b := report[i], report[i+1]

		if i == dampened {
			a = report[i-1]
		}
		if i+1 == dampened {
			b = report[i+2]
		}

		diff := b - a

		switch direction {
		case 0:
			{
				if diff > 0 {
					direction = 1
				}
				if diff < 0 {
					direction = -1
				}
			}
		case 1:
			if diff < 0 {
				return false
			}
		case -1:
			if diff > 0 {
				return false
			}

		}

		absDiff := abs(diff)
		if absDiff <= 0 || absDiff > 3 {
			return false
		}

	}

	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

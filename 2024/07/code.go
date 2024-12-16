package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type equation struct {
	test int
	xs   []int
}

func run(part2 bool, input string) any {

	equations := []equation{}

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		split := strings.Split(line, ":")
		equation := equation{parseInt(split[0]), []int{}}

		for _, input := range strings.Split(strings.TrimSpace(split[1]), " ") {
			equation.xs = append(equation.xs, parseInt(input))
		}

		equations = append(equations, equation)
	}

	totalCalibrationResult := 0

	for _, equation := range equations {
		running, xs := equation.xs[0], equation.xs[1:]
		if part2 {
			if canEqualTest(equation.test, running, xs, []string{"+", "*", "||"}) {
				totalCalibrationResult += equation.test
			}
		} else {
			if canEqualTest(equation.test, running, xs, []string{"+", "*"}) {
				totalCalibrationResult += equation.test
			}
		}
	}

	return totalCalibrationResult
}

func canEqualTest(test, running int, xs []int, operations []string) bool {

	if running > test {
		return false
	}

	if len(xs) == 0 {
		return running == test
	}

	x, xs := xs[0], xs[1:]

	for _, operation := range operations {
		var nextRunning int
		switch operation {
		case "+":
			nextRunning = running + x
		case "*":
			nextRunning = running * x
		case "||":
			nextRunning = concat(running, x)
		default:
			panic("unknown operation")
		}

		if canEqualTest(test, nextRunning, xs, operations) {
			return true
		}
	}

	return false
}

func concat(x, y int) int {
	return parseInt(strconv.Itoa(x) + strconv.Itoa(y))
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

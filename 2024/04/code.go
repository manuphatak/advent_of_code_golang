package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}

	lines := strings.Split(strings.TrimSpace(input), "\n")

	matches := 0
	for i, line := range lines {
		for j := range line {
			matches += countMatches(lines, i, j)
		}
	}
	return matches
}

type direction struct {
	x, y int
}

var directions = [8]direction{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

func countMatches(lines []string, i, j int) int {
	count := 0
	for _, direction := range directions {
		if lines[i][j] != 'X' {
			continue
		}

		if !lineInbound(lines, i+(direction.y*3)) {
			continue
		}
		if !rowInbound(lines, j+(direction.x*3)) {
			continue
		}

		if lines[i+(direction.y*1)][j+(direction.x*1)] != 'M' {
			continue
		}
		if lines[i+(direction.y*2)][j+(direction.x*2)] != 'A' {
			continue
		}
		if lines[i+(direction.y*3)][j+(direction.x*3)] != 'S' {
			continue
		}
		count += 1
	}

	return count
}

func lineInbound(lines []string, i int) bool {
	return i >= 0 && i < len(lines)
}

func rowInbound(lines []string, j int) bool {
	return j >= 0 && j < len(lines[0])
}

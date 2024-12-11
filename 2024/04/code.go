package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type direction struct {
	x, y int
}

// part 1 directions
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

func run(part2 bool, input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	matches := 0
	for i, line := range lines {
		for j := range line {
			if part2 {
				matches += countMatchesPart2(lines, i, j)
			} else {
				matches += countMatchesPart1(lines, i, j)
			}
		}
	}
	return matches
}

func countMatchesPart1(lines []string, i, j int) int {
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

func countMatchesPart2(lines []string, i, j int) int {
	count := 0

	if (masDR(lines, i, j) || samDR(lines, i, j)) && (masDL(lines, i, j+2) || samDL(lines, i, j+2)) {
		count += 1
	}

	return count
}

func masDR(lines []string, i int, j int) bool {
	if !lineInbound(lines, i+2) || !rowInbound(lines, j+2) {
		return false
	}
	return lines[i][j] == 'M' && lines[i+1][j+1] == 'A' && lines[i+2][j+2] == 'S'
}
func samDR(lines []string, i int, j int) bool {
	if !lineInbound(lines, i+2) || !rowInbound(lines, j+2) {
		return false
	}
	return lines[i][j] == 'S' && lines[i+1][j+1] == 'A' && lines[i+2][j+2] == 'M'
}
func masDL(lines []string, i int, j int) bool {
	if !lineInbound(lines, i+2) || !rowInbound(lines, j-2) {
		return false
	}
	return lines[i][j] == 'M' && lines[i+1][j-1] == 'A' && lines[i+2][j-2] == 'S'
}
func samDL(lines []string, i int, j int) bool {
	if !lineInbound(lines, i+2) || !rowInbound(lines, j-2) {
		return false
	}
	return lines[i][j] == 'S' && lines[i+1][j-1] == 'A' && lines[i+2][j-2] == 'M'
}

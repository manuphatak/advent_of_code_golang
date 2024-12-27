package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type point struct {
	y, x int
}

func isInBounds(tm [][]int, p point) bool {
	return p.y >= 0 && p.y < len(tm) && p.x >= 0 && p.x < len(tm[0])
}

type Set[E comparable] map[E]struct{}

func NewSet[E comparable]() Set[E] { return make(Set[E]) }
func (s Set[E]) Add(v E) {
	s[v] = struct{}{}
}

func (s Set[E]) Contains(v E) bool {
	_, ok := s[v]
	return ok
}

func (s Set[E]) String() string {
	keys := []string{}
	for k := range s {
		keys = append(keys, fmt.Sprintf("%v", k))
	}

	return fmt.Sprintf("Set[%v]", strings.Join(keys, " "))
}

func (s Set[E]) Size() int {
	return len(s)
}

func run(part2 bool, input string) any {
	trailMap := [][]int{}
	trailHeads := []point{}

	for i, row := range strings.Split(strings.TrimSpace(input), "\n") {
		trailMapRow := []int{}
		for j, cell := range row {
			trailMapRow = append(trailMapRow, parseInt(string(cell)))
			if cell == '0' {
				trailHeads = append(trailHeads, point{i, j})
			}
		}

		trailMap = append(trailMap, trailMapRow)
	}

	score, rating := 0, 0

	for _, trailHead := range trailHeads {
		trailPeaks := NewSet[point]()
		rating += walk(trailMap, trailHead, trailPeaks, 1)
		score += trailPeaks.Size()
	}
	if part2 {
		return rating
	} else {
		return score
	}
}

// walk traverses the trailMap starting from the given point `p`, looking for
// adjacent points that match the `matchValue`. It uses a depth-first search
// approach to explore all valid paths. The function updates the `trailPeaks`
// set with points that have a value of 9 and returns a rating based on the
// number of such peaks found.
//
// Parameters:
// - trailMap: A 2D slice of integers representing the trail map.
// - p: The starting point for the traversal.
// - trailPeaks: A set to store points that have a value of 9.
// - matchValue: The value to match in the trail map.
//
// Returns:
// - An integer rating based on the number of peaks found.
func walk(trailMap [][]int, p point, trailPeaks Set[point], matchValue int) int {
	var directions = [...]point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	rating := 0

	for _, direction := range directions {
		nextPoint := point{p.y + direction.y, p.x + direction.x}
		if !isInBounds(trailMap, nextPoint) {
			continue
		}

		nextValue := trailMap[nextPoint.y][nextPoint.x]
		if nextValue != matchValue {
			continue
		}

		if nextValue == 9 {
			trailPeaks.Add(nextPoint)
			rating++
			continue
		}
		rating += walk(trailMap, nextPoint, trailPeaks, matchValue+1)
	}

	return rating
}

func parseInt(s string) int {
	if s == "." {
		return -1
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

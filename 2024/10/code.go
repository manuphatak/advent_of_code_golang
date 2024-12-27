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

func IsInBounds(tm [][]int, p point) bool {
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
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}

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

	score := 0
	for _, trailHead := range trailHeads {
		trailPeaks := NewSet[point]()
		walk(trailMap, trailHead, trailPeaks, 1)
		score += trailPeaks.Size()
	}

	return score
}

func walk(trailMap [][]int, p point, trailPeaks Set[point], matchValue int) {
	var directions = [...]point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for _, direction := range directions {
		nextPoint := point{p.y + direction.y, p.x + direction.x}
		if !IsInBounds(trailMap, nextPoint) {
			continue
		}

		nextValue := trailMap[nextPoint.y][nextPoint.x]
		if nextValue != matchValue {
			continue
		}

		if nextValue == 9 {
			trailPeaks.Add(nextPoint)
			continue
		}
		walk(trailMap, nextPoint, trailPeaks, matchValue+1)
	}
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

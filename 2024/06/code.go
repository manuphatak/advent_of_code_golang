package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type direction int

const (
	UP direction = iota
	RIGHT
	DOWN
	LEFT
)

type point struct {
	x, y int
}

type guard struct {
	point point
	dir   direction
}

func (g guard) IsOnMap(cols, rows int) bool {
	return g.point.x >= 0 && g.point.x < cols && g.point.y >= 0 && g.point.y < rows

}

func (g guard) Next() point {

	switch g.dir {
	case UP:
		return point{g.point.x, g.point.y - 1}
	case RIGHT:
		return point{g.point.x + 1, g.point.y}
	case DOWN:
		return point{g.point.x, g.point.y + 1}
	case LEFT:
		return point{g.point.x - 1, g.point.y}

	}

	panic("unknown direction")
}

type state struct {
	obstacles  pointSet
	visited    pointSet
	guard      guard
	rows, cols int
}

func newState(rows int, cols int) *state {
	return &state{
		obstacles: pointSet{},
		visited:   pointSet{},
		guard:     guard{},
		rows:      rows,
		cols:      cols,
	}
}

func (s *state) SetGuard(guard guard) {
	s.guard = guard
}

func (s state) GuardIsOnMap() bool {
	return s.guard.IsOnMap(s.cols, s.rows)
}
func (s *state) MarkPositionVisited() {
	s.visited.Add(s.guard.point)

}
func (s state) ObstacleInNext() bool {
	return s.obstacles.Has(s.guard.Next())
}
func (s *state) TurnRight() {
	switch s.guard.dir {
	case UP:
		s.guard.dir = RIGHT
	case RIGHT:
		s.guard.dir = DOWN
	case DOWN:
		s.guard.dir = LEFT
	case LEFT:
		s.guard.dir = UP
	}

}
func (s *state) MoveForward() {
	s.guard.point = s.guard.Next()
}

type pointSet map[point]struct{}

func (s pointSet) Add(p point) {
	s[p] = struct{}{}
}
func (s pointSet) Has(p point) bool {
	_, ok := s[p]
	return ok
}

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}

	// 	read in the map, find the guard position
	lines := strings.Split(strings.TrimSpace(input), "\n")

	state := newState(len(lines), len(lines[0]))

	for i, line := range lines {
		for j, cell := range line {
			if cell == '#' {
				state.obstacles.Add(point{j, i})
			}
			if cell == '^' {
				state.SetGuard(guard{point{j, i}, UP})

			}
		}
	}

	for {
		if state.GuardIsOnMap() {
			state.MarkPositionVisited()
		} else {
			break
		}

		if state.ObstacleInNext() {
			state.TurnRight()
		} else {
			state.MoveForward()
		}
	}

	return len(state.visited)
}

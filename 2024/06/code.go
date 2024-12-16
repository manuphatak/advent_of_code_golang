package main

import (
	"maps"
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
	obstacles     pointSet
	visited       pointSet
	loopDetection guardSet
	guard         guard
	rows, cols    int
}

func newState(rows int, cols int) *state {
	return &state{
		obstacles:     pointSet{},
		visited:       pointSet{},
		loopDetection: guardSet{},
		guard:         guard{},
		rows:          rows,
		cols:          cols,
	}
}

func (s *state) Run() (int, bool) {
	for {
		if s.guardIsOnMap() {
			s.markPositionVisited()
		} else {
			break
		}

		if s.obstacleInNext() {
			ok := s.turnRight()

			if !ok {
				return len(s.visited), false
			}
		} else {
			s.moveForward()
		}
	}

	return len(s.visited), true
}

func (s *state) SetGuard(guard guard) {
	s.guard = guard
}
func (s state) Clone() *state {
	nextState := newState(s.rows, s.cols)
	nextState.obstacles = maps.Clone(s.obstacles)
	nextState.SetGuard(s.guard)

	return nextState
}

func (s state) guardIsOnMap() bool {
	return s.guard.IsOnMap(s.cols, s.rows)
}
func (s *state) markPositionVisited() {
	s.visited.Add(s.guard.point)

}
func (s state) obstacleInNext() bool {
	return s.obstacles.Has(s.guard.Next())
}
func (s *state) turnRight() bool {
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

	if s.loopDetection.Has(s.guard) {
		return false
	} else {
		s.loopDetection.Add(s.guard)
		return true
	}

}
func (s *state) moveForward() {
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

type guardSet map[guard]struct{}

func (s guardSet) Add(g guard) {
	s[g] = struct{}{}
}
func (s guardSet) Has(g guard) bool {
	_, ok := s[g]
	return ok
}

func run(part2 bool, input string) any {

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

	if part2 {
		count := 0
		for i := 0; i < state.rows; i++ {
			for j := 0; j < state.cols; j++ {
				// for each point...
				point := point{j, i}

				// ...if it's an obstacle, skip it
				if state.obstacles.Has(point) {
					continue
				}

				// ...otherwise, clone the state, add the obstacle, and run the simulation
				newState := state.Clone()
				newState.obstacles.Add(point)
				_, ok := newState.Run()

				// if the simulation failed (found an infinite loop), increment the count
				if !ok {
					count += 1
				}

			}
		}

		return count

	} else {
		state.obstacles.Add(point{0, 0})
		visited, ok := state.Run()

		if !ok {
			panic("loop detected in part 1")
		}
		return visited
	}
}

package main

import (
	"iter"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type point struct {
	y, x int
}

type Set[K comparable] map[K]struct{}

func NewSet[K comparable]() Set[K] {
	return Set[K]{}
}

func (s Set[K]) Add(v K) {
	s[v] = struct{}{}
}

func run(part2 bool, input string) any {
	antennaFrequencies := map[string][]point{}

	lines := strings.Split(strings.TrimSpace(input), "\n")
	rows, cols := len(lines), len(lines[0])

	for y, line := range lines {
		for x, cell := range line {
			if cell == '.' {
				continue
			}

			frequency := string(cell)
			antennaFrequencies[frequency] = append(antennaFrequencies[frequency], point{y, x})
		}
	}

	antinodes := Set[point]{}

	for _, antennas := range antennaFrequencies {
		for pair := range permutations2(antennas) {
			distance := sub(pair[1], pair[0])

			if part2 {
				addAllOnPath := func(node point, move func(node, distance point) point) {
					antinodes.Add(node)
					for {
						node = move(node, distance)
						if inRange(node, rows, cols) {
							antinodes.Add(node)
						} else {
							break
						}
					}
				}

				addAllOnPath(pair[0], sub)
				addAllOnPath(pair[1], add)
			} else {
				if node := sub(pair[0], distance); inRange(node, rows, cols) {
					antinodes.Add(node)
				}

				if node := add(pair[1], distance); inRange(node, rows, cols) {
					antinodes.Add(node)
				}

			}
		}
	}

	return len(antinodes)
}

func inRange(node point, rows, cols int) bool {
	return node.x >= 0 && node.x < cols && node.y >= 0 && node.y < rows
}

func add(l, r point) point {
	return point{x: l.x + r.x, y: l.y + r.y}
}

func sub(l, r point) point {
	return point{x: l.x - r.x, y: l.y - r.y}
}

func permutations2(antennas []point) iter.Seq[[2]point] {
	return func(yield func([2]point) bool) {
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				yield([2]point{antennas[i], antennas[j]})
			}
		}
	}
}

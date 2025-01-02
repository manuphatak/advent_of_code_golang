package main

import (
	"iter"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
	"github.com/manuphatak/advent_of_code_golang/2024/shared"
)

func main() {
	aoc.Harness(run)
}

type Point struct {
	y, x int
}

type RegionTotal struct {
	area, fences int
}

func Pop(s shared.Set[Point]) (Point, bool) {
	for point := range s {
		s.Remove(point)
		return point, true
	}
	return Point{}, false
}

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}

	plots := make(map[Point]rune)
	queue := shared.NewSet[Point]()

	for y, row := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, char := range row {
			plots[Point{y, x}] = char
			queue.Add(Point{y, x})
		}
	}

	regionTotals := []RegionTotal{}
	for pos, ok := Pop(queue); ok; pos, ok = Pop(queue) {
		regionTotal := walkRegion(RegionTotal{}, plots, pos, queue, shared.NewSet[Point]())
		regionTotals = append(regionTotals, regionTotal)
	}

	return calculatePrice(regionTotals)
}

func calculatePrice(regionTotals []RegionTotal) int {
	price := 0

	for _, regionTotal := range regionTotals {
		price += (regionTotal.area * regionTotal.fences)
	}

	return price
}

func walkRegion(regionTotal RegionTotal, plots map[Point]rune, pos Point, queue, discoveredRegion shared.Set[Point]) RegionTotal {
	regionTotal.area++
	regionTotal.fences += 4
	for adjPos := range adjacent(pos) {
		if plots[adjPos] == plots[pos] {
			regionTotal.fences--
			discoveredRegion.Add(adjPos)
		}
	}

	if pos, ok := findNext(queue, discoveredRegion); ok {
		return walkRegion(regionTotal, plots, pos, queue, discoveredRegion)
	}

	return regionTotal
}

func findNext(queue, discoveredRegion shared.Set[Point]) (Point, bool) {
	pos, ok := Pop(discoveredRegion)

	if !ok {
		return pos, false
	}

	if queue.Contains(pos) {
		queue.Remove(pos)
		return pos, true

	}
	return findNext(queue, discoveredRegion)
}

func adjacent(point Point) iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for _, dir := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			if ok := yield(Point{point.y + dir.y, point.x + dir.x}); !ok {
				break
			}
		}
	}
}

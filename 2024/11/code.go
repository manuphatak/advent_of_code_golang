package main

import (
	"container/list"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}

	stones := list.New()

	for _, v := range strings.Split(strings.TrimSpace(input), " ") {
		stones.PushBack(parseInt(v))
	}

	blinks := 25
	for i := 0; i < blinks; i++ {
		for e := stones.Front(); e != nil; e = e.Next() {
			value := e.Value.(int)

			if value == 0 {
				e.Value = 1
			} else if isEven, l, r := splitDigits(value); isEven {
				stones.InsertBefore(l, e)
				e.Value = r
			} else {
				e.Value = value * 2024
			}
		}
	}

	// solve part 1 here
	return stones.Len()
}

func splitDigits(value int) (bool, int, int) {
	if value < 10 {
		return false, -1, -1
	}
	stringValue := strconv.Itoa(value)
	isEven := len(stringValue)%2 == 0
	midPoint := len(stringValue) / 2
	return isEven, parseInt(stringValue[:midPoint]), parseInt(stringValue[midPoint:])

}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

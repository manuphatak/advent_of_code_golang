package main

import (
	"strconv"
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

	parsedInput := []int{}

	for _, char := range strings.TrimSpace(input) {
		parsedInput = append(parsedInput, parseInt(string(char)))
	}

	filesystem := compact(parsedInput)

	return checksum(filesystem)
}

func checksum(filesystem []int) int {
	sum := 0

	for i, n := range filesystem {
		sum += i * n
	}

	return sum
}

func compact(input []int) []int {
	filesystemSize := 0
	for i, n := range input {
		if i%2 == 0 {
			filesystemSize += n
		}
	}

	filesystem := make([]int, filesystemSize)
	pos := 0
	freeSpaceQueue := make(chan int, filesystemSize)
	overflowStack := []int{}

	for i := 0; i < len(input); i++ {
		n := input[i]
		id := i / 2
		isFileBlock := i%2 == 0

		for j := 0; j < n; j++ {
			// While there are still blocks to write
			if pos < filesystemSize {
				if isFileBlock {
					filesystem[pos] = id
					pos++
				} else {
					freeSpaceQueue <- pos
					pos++
				}
			} else {
				// Otherwise, store the block in the overflow stack
				if isFileBlock {
					overflowStack = append(overflowStack, id)
				}
			}
		}
	}

	for i := len(overflowStack) - 1; i >= 0; i-- {
		n := overflowStack[i]
		freeSlot := <-freeSpaceQueue
		filesystem[freeSlot] = n
	}

	return filesystem
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

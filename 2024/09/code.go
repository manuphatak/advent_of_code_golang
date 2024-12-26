package main

import (
	"container/list"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	parsedInput := []int{}

	for _, char := range strings.TrimSpace(input) {
		parsedInput = append(parsedInput, parseInt(string(char)))
	}

	var filesystem []int
	if part2 {
		filesystem = compactDefragment(parsedInput)
	} else {
		filesystem = compactFragment(parsedInput)
	}

	return checksum(filesystem)
}

func checksum(filesystem []int) int {
	sum := 0

	for i, n := range filesystem {
		sum += i * n
	}

	return sum
}

func compactFragment(input []int) []int {
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

	for i, n := range input {
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

type block interface {
	CanFit(block) bool
	Size() int
	Equals(block) bool
}

type fileBlock struct {
	id, size int
}

func (f fileBlock) String() string {
	return fmt.Sprintf("file{%d %d}", f.id, f.size)
}
func (f fileBlock) Size() int {
	return f.size
}

func (f fileBlock) CanFit(b block) bool {
	return false
}
func (f fileBlock) Equals(b block) bool {

	switch b := b.(type) {
	case fileBlock:
		return b.id == f.id
	default:
		return false
	}

}

type freeBlock struct {
	size int
}

func (f freeBlock) String() string {
	return fmt.Sprintf("free{_ %d}", f.size)
}

func (f freeBlock) Size() int {
	return f.size
}

func (f freeBlock) CanFit(b block) bool {
	return f.Size() >= b.Size()
}
func (f freeBlock) Equals(b block) bool {
	return false
}

func compactDefragment(input []int) []int {
	filesystem := list.New()
	fileBlockElements := []*list.Element{}

	for i, size := range input {
		id := i / 2
		isFileBlock := i%2 == 0

		if isFileBlock {
			fileBlockElements = append(fileBlockElements,
				filesystem.PushBack(fileBlock{id, size}))
		} else {
			filesystem.PushBack(freeBlock{size})
		}
	}

	slices.Reverse(fileBlockElements)
	for _, e := range fileBlockElements {
		for mark := filesystem.Front(); mark != nil; mark = mark.Next() {
			markValue, eValue := mark.Value.(block), e.Value.(block)
			if markValue.Equals(eValue) {
				// in order to avoid moving the file block backwards
				break
			}
			if markValue.CanFit(eValue) {
				// insert a free block that will eventually replace the file
				insertedSpace := filesystem.InsertBefore(freeBlock{eValue.Size()}, e)

				// move the file block to the free space (free space to be removed)
				filesystem.MoveBefore(e, mark)
				// if there is remaining space, insert a free block
				if remainingSize := markValue.Size() - eValue.Size(); remainingSize > 0 {
					filesystem.InsertBefore(freeBlock{remainingSize}, mark)
				}
				filesystem.Remove(mark)
				squashFreeBlocks(insertedSpace, filesystem)

				break
			}
		}
	}

	output := []int{}

	for e := filesystem.Front(); e != nil; e = e.Next() {
		for i := 0; i < e.Value.(block).Size(); i++ {
			switch e.Value.(type) {
			case fileBlock:
				output = append(output, e.Value.(fileBlock).id)
			case freeBlock:
				output = append(output, 0)
			}

		}
	}

	return output

}

func squashFreeBlocks(insertedSpace *list.Element, filesystem *list.List) {
	spaceToAdd := 0
	if insertedSpace.Prev() != nil {
		switch insertedSpace.Prev().Value.(type) {
		case freeBlock:
			spaceToAdd += filesystem.Remove(insertedSpace.Prev()).(freeBlock).Size()
		}
	}
	if insertedSpace.Next() != nil {
		switch insertedSpace.Next().Value.(type) {
		case freeBlock:
			spaceToAdd += filesystem.Remove(insertedSpace.Next()).(freeBlock).Size()
		}
	}

	if spaceToAdd > 0 {
		filesystem.InsertBefore(freeBlock{spaceToAdd + insertedSpace.Value.(block).Size()}, insertedSpace)
		filesystem.Remove(insertedSpace)
	}
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

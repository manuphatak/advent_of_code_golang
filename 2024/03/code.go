package main

import (
	"regexp"
	"strconv"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Command interface {
	isCommand()
}
type Mul struct {
	l int
	r int
}

func (m Mul) isCommand() {}

type Do struct{}

func (d Do) isCommand() {}

type Dont struct{}

func (d Dont) isCommand() {}

func run(part2 bool, input string) any {

	commands := parseCommands(input)

	sumProduct := 0

	if part2 {
		enabled := true
		for _, command := range commands {
			switch c := command.(type) {
			case Mul:
				if enabled {
					sumProduct += c.l * c.r
				}
			case Do:
				enabled = true
			case Dont:
				enabled = false
			}
		}
	} else {
		for _, command := range commands {
			switch c := command.(type) {
			case Mul:
				sumProduct += c.l * c.r
			}
		}
	}

	return sumProduct
}

func parseCommands(input string) []Command {
	re := regexp.MustCompile(`(?P<mul>mul\((?P<l>\d{0,3}),(?P<r>\d{0,3})\))|(?P<do>do\(\))|(?P<dont>don't\(\))`)

	mulIndex, lIndex, rIndex, doIndex, dontIndex := re.SubexpIndex("mul"), re.SubexpIndex("l"), re.SubexpIndex("r"), re.SubexpIndex("do"), re.SubexpIndex("dont")

	commands := []Command{}
	for _, match := range re.FindAllStringSubmatch(input, -1) {
		if match[mulIndex] != "" {
			commands = append(commands, Mul{parseInt(match[lIndex]), parseInt(match[rIndex])})
		} else if match[doIndex] != "" {
			commands = append(commands, Do{})
		} else if match[dontIndex] != "" {
			commands = append(commands, Dont{})
		}
	}
	return commands
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

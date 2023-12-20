package main

import (
	"fmt"
	"strings"

	"github.com/patricho/advent-of-code/go/util"
)

type Map struct {
	Left  string
	Right string
}

func main() {
	step1("test1.txt")
	step1("test2.txt")

	util.Measure(func() {
		step1("input.txt")
	})

	step2("test3.txt")

	util.Measure(func() {
		step2("input.txt")
	})
}

func step1(filename string) {
	lines := util.ReadFile(filename)
	directions := lines[0]
	mapping := parseMapping(lines)
	steps := find(mapping, directions, "AAA", "ZZZ")
	fmt.Println("steps: ", steps)
}

func step2(filename string) {
	lines := util.ReadFile(filename)
	directions := lines[0]
	mapping := parseMapping(lines)
	steps := []int{}

	// find all starting points **A
	for key := range mapping {
		if !strings.HasSuffix(key, "A") {
			continue
		}

		// exit on the first found **Z
		steps = append(steps, find(mapping, directions, key, "Z"))
	}

	// calculate LCM for all the found steps
	fmt.Println("found steps:", steps)
	fmt.Println("steps lcm:", util.LCM(steps...))
}

func parseMapping(lines []string) map[string]Map {
	mapping := map[string]Map{}

	for i := 2; i < len(lines); i++ {
		arr := strings.Split(strings.TrimSuffix(lines[i], ")"), " = (")
		key := arr[0]
		vals := strings.Split(arr[1], ", ")
		mapping[key] = Map{
			Left:  vals[0],
			Right: vals[1],
		}
	}

	return mapping
}

func find(mapping map[string]Map, directions, start, find string) int {
	key := start
	step := 0

	for {
		step++
		dir := string(directions[(step-1)%len(directions)])
		curr := mapping[key]
		var next string
		if dir == "L" {
			next = curr.Left
		} else {
			next = curr.Right
		}

		// fmt.Println("step", step, "curr", key, "dir", dir, "next", next)

		if strings.HasSuffix(next, find) {
			break
		} else {
			key = next
		}
	}

	return step
}

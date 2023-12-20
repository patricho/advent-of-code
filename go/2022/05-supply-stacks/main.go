package supplystacks05

import (
	"fmt"
	"regexp"
	"strconv"

	c "github.com/patricho/advent-of-code/go/util"
)

/*
test-input:

    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2

expected result: CMZ
*/

func Part1() {
	stacks := populateStacks()

	// instructions
	instructions := c.ReadFile("05-supply-stacks/input.txt")

	for _, instr := range instructions {
		num, fromidx, toidx := parseInstructions(instr)

		for i := 1; i <= num; i++ {
			move(&stacks[fromidx], &stacks[toidx])
		}
	}

	answer := []string{}
	for i := 0; i < len(stacks); i++ {
		answer = append(answer, stacks[i][0])
	}

	fmt.Println("answer", answer)
}

func Part2() {
	stacks := populateStacks()

	for i, stack := range stacks {
		fmt.Println("stack", i, stack)
	}

	// instructions
	instructions := c.ReadFile("05-supply-stacks/input.txt")

	for _, instr := range instructions {
		num, fromidx, toidx := parseInstructions(instr)
		moveMultiple(num, &stacks[fromidx], &stacks[toidx])
	}

	answer := []string{}
	for i := 0; i < len(stacks); i++ {
		answer = append(answer, stacks[i][0])
	}

	fmt.Println("answer", answer)
}

func populateStacks() [][]string {
	stacks := make([][]string, 9)
	idxstart, idxoffset := 1, 4

	stackslines := c.ReadFile("05-supply-stacks/stacks.txt")

	for _, line := range stackslines {
		fmt.Println("stackslines", line)

		for i := 0; i < 9; i++ {
			idx := idxstart + idxoffset*i
			char := string(line[idx])

			if char == " " {
				continue
			}

			stacks[i] = append(stacks[i], string(line[idx]))
		}
	}

	for i := 0; i < 9; i++ {
		fmt.Println("stacks", i, stacks[i])
	}

	return stacks
}

func populateTestStacks() [][]string {
	// starting stacks for test
	stacks := make([][]string, 3)

	stacks[0] = []string{"N", "Z"}
	stacks[1] = []string{"D", "C", "M"}
	stacks[2] = []string{"P"}

	return stacks
}

func parseInstructions(input string) (num, fromidx, toidx int) {
	re, _ := regexp.Compile(`^move (\d+) from (\d) to (\d)$`)

	num, _ = strconv.Atoi(re.ReplaceAllString(input, "$1"))
	fromidx, _ = strconv.Atoi(re.ReplaceAllString(input, "$2"))
	toidx, _ = strconv.Atoi(re.ReplaceAllString(input, "$3"))

	return num, fromidx - 1, toidx - 1
}

func move(from *[]string, to *[]string) {
	// "move 1 from 2 to 1"

	// select from stack 2
	box := (*from)[0]

	// remove from 2
	*from = (*from)[1:]

	// add to 1
	*to = append(*to, "")      // add empty element at end
	copy((*to)[1:], (*to)[0:]) // shift down from index 0 to index 1 (overwrites last element, leaves index 0 available)
	(*to)[0] = box             // add selected item to beginning of stack
}

func moveMultiple(num int, from *[]string, to *[]string) {
	// "move 2 from 2 to 1"

	// select n boxes from stack 2
	var boxes []string

	// [0:0] gives empty result?
	if num > 1 {
		boxes = (*from)[0:num]
	} else {
		boxes = (*from)[0:1]
	}

	// remove n boxes from 2
	*from = (*from)[num:]

	// add n boxes to 1
	*to = append(*to, boxes...)  // add n empty elements at end
	copy((*to)[num:], (*to)[0:]) // shift down from index 0 to index n (overwrites last elements, leaves index 0-(n-1) available)

	copy((*to)[0:], boxes) // add selected items to beginning of stack
}

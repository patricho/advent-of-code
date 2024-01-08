package main

import (
	"fmt"
	"slices"

	s "github.com/patricho/advent-of-code/go/shared"
)

var (
	Add  int = 1
	Mult int = 2
	Halt int = 99
)

func main() {
	// runProgram("1,9,10,3,2,3,11,0,99,30,40,50")
	// runProgram("1,0,0,0,99")
	// runProgram("2,3,0,3,99")
	// runProgram("2,4,4,5,99,0")
	// runProgram("1,1,1,4,99,5,6,0,99")

	input := "1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,19,9,23,1,5,23,27,1,27,9,31,1,6,31,35,2,35,9,39,1,39,6,43,2,9,43,47,1,47,6,51,2,51,9,55,1,5,55,59,2,59,6,63,1,9,63,67,1,67,10,71,1,71,13,75,2,13,75,79,1,6,79,83,2,9,83,87,1,87,6,91,2,10,91,95,2,13,95,99,1,9,99,103,1,5,103,107,2,9,107,111,1,111,5,115,1,115,5,119,1,10,119,123,1,13,123,127,1,2,127,131,1,131,13,0,99,2,14,0,0"

	s.Measure(func() {
		part1 := parseAndRunProgram(input)
		fmt.Println("part 1", part1)
	})
	s.Measure(func() {
		part2 := iteratePrograms(input, 19690720)
		fmt.Println("part 2", part2)
	})
}

func iteratePrograms(input string, target int) int {
	startcodes := s.ToIntSlice(input)
	length := len(startcodes)
	sum, one, two := 0, 0, 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			codes := slices.Clone(startcodes)
			codes[1] = i
			codes[2] = j
			sum = runProgram(codes)
			if sum == target {
				one = i
				two = j
				return 100*one + two
			}
		}
	}
	return 0
}

func parseAndRunProgram(input string) int {
	codes := s.ToIntSlice(input)
	return runProgram(codes)
}

func runProgram(codes []int) int {
	length := len(codes)

	for start := 0; start < length; start += 4 {
		op := codes[start]
		if op == Halt {
			break
		}

		sum := 0
		one := codes[start+1]
		two := codes[start+2]
		target := codes[start+3]

		if one >= length || two >= length || target >= length {
			continue
		}

		if op == Add {
			sum = codes[one] + codes[two]
		} else if op == Mult {
			sum = codes[one] * codes[two]
		}
		codes[target] = sum
	}

	return codes[0]
}

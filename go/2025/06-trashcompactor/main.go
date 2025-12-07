package main

import (
	"strings"

	s "github.com/patricho/advent-of-code/go/shared"
)

type Equation struct {
	Numbers []int
	Sign    rune
}

const (
	INPUT_FILE = "../inputs/2025/06-input.txt"
	TEST_FILE  = "../inputs/2025/06-test.txt"
)

func main() {
	s.DisplayMain(func() {
		test()
		run()
	})
}

func test() {
	s.RunCase("test part 1", func() int { return part1(TEST_FILE) }, 4277556)
	s.RunCase("test part 2", func() int { return part2(TEST_FILE) }, 3263827)
}

func run() {
	s.RunCase("part 1", func() int { return part1(INPUT_FILE) }, 5060053676136)
	s.RunCase("part 2", func() int { return part2(INPUT_FILE) }, 9695042567249)
}

func part1(filename string) int {
	eqs := parseInput(filename)
	return sumEquations(eqs)
}

func part2(filename string) int {
	lines := s.ReadFile(filename)
	columns := len(lines[0])
	eqs := []Equation{}
	pending := Equation{}

	// Check column for column from the end
	for x := columns - 1; x >= 0; x-- {
		empty := true
		numstring := ""

		// Create a number from the strings top to bottom
		for _, line := range lines {
			n := line[x]
			if n == ' ' {
				continue
			}

			// Remember the sign if we see it, otherwise add to the number
			if n == '*' || n == '+' {
				pending.Sign = rune(n)
			} else {
				numstring += string(n)
			}
			empty = false
		}

		if empty {
			// When we encounter an empty column, the equation is done, so add it to the list
			eqs = append(eqs, pending)
			pending = Equation{}
		} else {
			// Otherwise, add the number
			pending.Numbers = append(pending.Numbers, s.ToInt(numstring))
		}
	}

	// Add the final equation too, from the first columns
	if len(pending.Numbers) > 0 {
		eqs = append(eqs, pending)
	}

	return sumEquations(eqs)
}

func parseInput(filename string) []Equation {
	lines := s.ReadFile(filename)

	lastline := lines[len(lines)-1]

	columns := strings.Fields(lastline)

	eqs := make([]Equation, len(columns))

	for i := 0; i < len(lines)-1; i++ {
		parts := strings.Fields(lines[i])
		numbers := make([]int, len(parts))
		for j := 0; j < len(parts); j++ {
			numbers[j] = s.ToInt(parts[j])
			eqs[j].Numbers = append(eqs[j].Numbers, numbers[j])
		}
	}

	lastLineParts := strings.Fields(lastline)

	for j := 0; j < len(columns); j++ {
		eqs[j].Sign = rune(lastLineParts[j][0])
	}

	return eqs
}

func sumEquations(eqs []Equation) int {
	result := 0

	for _, eq := range eqs {
		res := eq.Numbers[0]
		switch eq.Sign {
		case '*':
			for i := 1; i < len(eq.Numbers); i++ {
				res *= eq.Numbers[i]
			}
		case '+':
			for i := 1; i < len(eq.Numbers); i++ {
				res += eq.Numbers[i]
			}
		}
		result += res
	}

	return result
}

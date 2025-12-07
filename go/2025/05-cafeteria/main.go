package main

import (
	"slices"
	"strings"

	s "github.com/patricho/advent-of-code/go/shared"
)

const (
	INPUT_FILE = "../inputs/2025/05-input.txt"
	TEST_FILE  = "../inputs/2025/05-test.txt"
)

func main() {
	s.DisplayMain(func() {
		test()
		run()
	})
}

func test() {
	s.RunCase("test part 1", func() int { return part1(TEST_FILE) }, 3)
	s.RunCase("test part 2", func() int { return part2(TEST_FILE) }, 14)
}

func run() {
	s.RunCase("part 1", func() int { return part1(INPUT_FILE) }, 598)
	s.RunCase("part 2", func() int { return part2(INPUT_FILE) }, 360341832208407)
}

func parseInput(filename string) ([]int, [][]int) {
	lines := s.ReadFile(filename)

	ingredients := []int{}
	ranges := [][]int{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, "-")
		if len(parts) == 2 {
			ranges = append(ranges, []int{s.ToInt(parts[0]), s.ToInt(parts[1])})
		} else {
			ingredients = append(ingredients, s.ToInt(line))
		}
	}

	return ingredients, ranges
}

func part1(filename string) int {
	result := 0

	ingredients, ranges := parseInput(filename)

	for _, ing := range ingredients {
		found := false
		for _, rng := range ranges {
			if ing >= rng[0] && ing <= rng[1] {
				found = true
				break
			}
		}
		if found {
			result++
		}
	}

	return result
}

func part2(filename string) int {
	result := 0

	_, ranges := parseInput(filename)

	slices.SortFunc(ranges, func(a []int, b []int) int {
		return a[0] - b[0]
	})

	maxCounted := -1

	for i := 0; i < len(ranges); i++ {
		cr := ranges[i]

		start := cr[0]
		end := cr[1]

		if maxCounted >= end {
			// Range is already completely covered
			continue
		}

		if maxCounted >= start {
			start = maxCounted + 1
		}

		span := end - start + 1

		result += span
		maxCounted = end
	}

	return result
}

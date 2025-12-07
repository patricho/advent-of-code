package main

import (
	s "github.com/patricho/advent-of-code/go/shared"
)

const (
	INPUT_FILE = "../inputs/2025/07-input.txt"
	TEST_FILE  = "../inputs/2025/07-test.txt"
)

var grid [][]rune

func main() {
	s.DisplayMain(func() {
		test()
		run()
	})
}

func test() {
	s.RunCase("test part 1", func() int { return part1(TEST_FILE) }, 21)
	s.RunCase("test part 2", func() int { return part2(TEST_FILE) }, 40)
}

func run() {
	s.RunCase("part 1", func() int { return part1(INPUT_FILE) }, 1658)
	s.RunCase("part 2", func() int { return part2(INPUT_FILE) }, 53916299384254)
}

func part1(filename string) int {
	splits, _ := traverse(filename)
	return splits
}

func part2(filename string) int {
	_, combinations := traverse(filename)
	return combinations
}

func initGrid(filename string) {
	lines := s.ReadFile(filename)
	grid = s.LinesToRuneGrid(lines)
}

func traverse(filename string) (int, int) {
	initGrid(filename)

	splits := 0
	beams := make([]int, len(grid[0]))

	// Find starting beam
	for i, r := range grid[0] {
		if r == 'S' {
			beams[i] = 1
			break
		}
	}

	for y := range grid {
		for x := range grid[y] {
			// We're only interested in when a beam reaches a splitter
			if grid[y][x] != '^' || beams[x] == 0 {
				continue
			}

			splits++

			if !s.OOByx(grid, y, x-1) {
				beams[x-1] += beams[x]
			}
			if !s.OOByx(grid, y, x+1) {
				beams[x+1] += beams[x]
			}

			beams[x] = 0
		}
	}

	combinations := 0
	for _, b := range beams {
		combinations += b
	}

	return splits, combinations
}

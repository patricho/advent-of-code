package main

import (
	s "github.com/patricho/advent-of-code/go/shared"
)

const (
	INPUT_FILE = "../inputs/2025/04-input.txt"
	TEST_FILE  = "../inputs/2025/04-test.txt"
)

var grid [][]rune

const (
	OFF  rune = '.'
	ON   rune = '@'
	MARK rune = 'X'
)

func main() {
	s.DisplayMain(func() {
		test()
		run()
	})
}

func test() {
	s.RunCase("test part 1", func() int { return part1(TEST_FILE) }, 13)
	s.RunCase("test part 2", func() int { return part2(TEST_FILE) }, 43)
}

func run() {
	s.RunCase("part 1", func() int { return part1(INPUT_FILE) }, 1491)
	s.RunCase("part 2", func() int { return part2(INPUT_FILE) }, 8722)
}

func part1(filename string) int {
	initGrid(filename)
	return countNeighbors()
}

func part2(filename string) int {
	result := 0

	initGrid(filename)

	for {
		round := countNeighbors()
		result += round

		if round == 0 {
			break
		}

		removeNeighbors()
	}

	return result
}

func initGrid(filename string) {
	lines := s.ReadFile(filename)
	grid = s.LinesToRuneGrid(lines)
}

func countNeighbors() int {
	result := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			pt := s.Point{X: x, Y: y}

			if grid[pt.Y][pt.X] == OFF {
				continue
			}

			neighbors := 0

			for _, dir := range s.DirectionsAndDiagonals {
				npt := s.Move(pt, dir)

				if s.OOB(grid, npt) {
					continue
				}

				if grid[npt.Y][npt.X] != OFF {
					neighbors++
				}
			}

			if neighbors < 4 {
				grid[pt.Y][pt.X] = MARK
				result++
			}
		}
	}

	// s.PrintGrid(grid, MARK)

	return result
}

func removeNeighbors() {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == MARK {
				grid[y][x] = OFF
			}
		}
	}
}

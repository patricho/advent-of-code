package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/patricho/advent-of-code/go/shared"
)

type Point struct {
	Char rune
	X    int
	Y    int
}

type CycleInfo struct {
	Sum     int
	Indexes []int
	Periods []int
}

var (
	grid   [][]rune
	width  int
	height int
	up     Point
	down   Point
	left   Point
	right  Point
	cycles map[string]CycleInfo
)

func main() {
	up = Point{Y: -1, X: 0}
	down = Point{Y: 1, X: 0}
	left = Point{Y: 0, X: -1}
	right = Point{Y: 0, X: 1}
	shared.Measure(func() {
		run(1, "input.txt")
	})
	shared.Measure(func() {
		run(2, "input.txt")
	})
}

func run(part int, filename string) {
	lines := shared.ReadFile(filename)
	grid = shared.LinesToRuneGrid(lines)
	width = len(grid[0])
	height = len(grid)
	cycles = map[string]CycleInfo{}

	// fmt.Println("start:")
	// printGrid(grid)
	// fmt.Println("")

	if part == 1 {
		sum := moveToSide(up, true)

		// fmt.Println("result:")
		// printGrid(grid)
		// fmt.Println("")

		fmt.Println("part 1 result:", sum)
		return
	}

	cycleNTimes(200)

	// analyze the first 200 cycles to find when it starts repeating, and the repeat period
	cycleStart := 200
	period := -1
	for _, c := range cycles {
		if c.Indexes[0] < cycleStart && len(c.Periods) > 0 {
			cycleStart = c.Indexes[0]
			period = c.Periods[0]
		}
	}

	// now do a roundabout calculation using the cached rows to find which cycle iteration 1 bn would land on
	// cycling starts at idx X, so start checking there
	// the target number we want is one billion, the target index is that -1
	// then check how many steps we need from X to reach target index
	stepsFromCycleStart := (1e9 - 1 - cycleStart) % period
	firstIdxToFind := cycleStart + stepsFromCycleStart

	// now find this row in the cached sums
	for hash, c := range cycles {
		if c.Indexes[0] != firstIdxToFind {
			continue
		}
		fmt.Println("part 2 result: cycle:", hash, ", start:", c.Indexes[0], ", period:", period, ", sum:", c.Sum)
	}
}

func cycleNTimes(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum = cycle()
		hash := shared.HashString(shared.GridToString(grid))
		c, ok := cycles[hash]
		if ok {
			if sum != c.Sum {
				panic("wtf")
			}
			period := i - c.Indexes[len(c.Indexes)-1]
			c.Indexes = append(c.Indexes, i)
			c.Periods = append(c.Periods, period)
			cycles[hash] = c
		} else {
			cycles[hash] = CycleInfo{
				Indexes: []int{i},
				Periods: []int{},
				Sum:     sum,
			}
		}
	}
	return sum
}

func cycle() int {
	moveToSide(up, true)
	moveToSide(left, false)
	moveToSide(down, false)
	sum := moveToSide(right, true)
	return sum
}

func moveToSide(dir Point, fromTopRight bool) int {
	if fromTopRight {
		// from top right to bottom left
		for y := 0; y < height; y++ {
			for x := width - 1; x >= 0; x-- {
				if grid[y][x] != 'O' {
					continue
				}
				rock := Point{Y: y, X: x}
				move(rock, dir)
			}
		}
	} else {
		// from bottom left to top right
		for y := height - 1; y >= 0; y-- {
			for x := 0; x < width; x++ {
				if grid[y][x] != 'O' {
					continue
				}
				rock := Point{Y: y, X: x}
				move(rock, dir)
			}
		}
	}
	return calculateWeight()
}

func calculateWeight() int {
	sum := 0
	for y := 0; y < height; y++ {
		rowidx := height - y
		rocks := 0
		for x := 0; x < width; x++ {
			if grid[y][x] == 'O' {
				rocks++
			}
		}
		sum += rowidx * rocks
	}
	return sum
}

func move(curr, dir Point) {
	next := curr
	for {
		next.X += dir.X
		next.Y += dir.Y
		if next.X < 0 || next.X >= width || next.Y < 0 || next.Y >= height { // out of bounds
			break
		} else if grid[next.Y][next.X] == '#' || grid[next.Y][next.X] == 'O' { // obstacle
			break
		}
		// if ok, do switch
		grid[next.Y][next.X] = 'O'
		grid[curr.Y][curr.X] = '.'
		curr.X += dir.X
		curr.Y += dir.Y
	}
}

func printGrid(grid [][]rune) {
	ce := color.New(color.FgRed)
	cf := color.New(color.Faint)
	for _, l := range grid {
		for _, r := range l {
			if r == 'O' {
				ce.Print("O")
			} else {
				cf.Print(string(r))
			}
		}
		fmt.Print("\n")
	}
}

package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/patricho/advent-of-code/go/shared"
)

type Point struct {
	X int
	Y int
}

var (
	lines       []string
	emptyRows   []int
	emptyCols   []int
	weightAddon int
)

func main() {
	// test
	run(2, "test.txt")
	run(10, "test.txt")
	run(100, "test.txt")

	// part 1
	shared.Measure(func() {
		run(2, "input.txt")
	})

	// part 2
	shared.Measure(func() {
		run(1_000_000, "input.txt")
	})
}

func run(weight int, filename string) {
	lines = shared.ReadFile(filename)
	weightAddon = weight - 1
	emptyRows, emptyCols = findEmpty()
	galaxies := findGalaxies()

	sum := 0

	for i := 0; i < len(galaxies); i++ {
		start := galaxies[i]

		for j := i + 1; j < len(galaxies); j++ {
			target := galaxies[j]

			steps := walk(start, target)
			sum += steps
		}
	}

	fmt.Println(filename, weightAddon+1, "total steps", sum)
}

func findGalaxies() []Point {
	hash := rune("#"[0])

	galaxies := []Point{}

	for lineidx, line := range lines {
		for charidx, char := range line {
			if char != hash {
				continue
			}
			galaxies = append(galaxies, Point{X: charidx, Y: lineidx})
		}
	}

	return galaxies
}

func walk(start, target Point) int {
	curr := start

	totalSteps := 0
	stepsOffset := 0

	for {
		dist := findDistance(curr, target)

		if atTarget(curr, target) {
			break
		}

		dir := findDirection(dist)
		curr, stepsOffset = takeStep(curr, dir)

		totalSteps += stepsOffset
	}

	return totalSteps
}

func takeStep(curr, dir Point) (next Point, steps int) {
	next = Point{
		X: curr.X + dir.X,
		Y: curr.Y + dir.Y,
	}

	weight := 1

	if emptyStep(next) {
		weight += weightAddon
	}

	return next, weight
}

func emptyStep(next Point) bool {
	return slices.Contains(emptyRows, next.Y) || slices.Contains(emptyCols, next.X)
}

func atTarget(curr, target Point) bool {
	return curr.X == target.X && curr.Y == target.Y
}

func findDirection(distance Point) Point {
	direction := Point{X: 0, Y: 0}

	// right or left
	if distance.X > 0 {
		direction.X = 1
	} else if distance.X < 0 {
		direction.X = -1
	}

	// down or up
	if distance.Y > 0 {
		direction.Y = 1
	} else if distance.Y < 0 {
		direction.Y = -1
	}

	// if there's movement to be done in both directions, start with just y
	if direction.X != 0 && direction.Y != 0 {
		direction.X = 0
	}

	return direction
}

func findDistance(curr, target Point) Point {
	return Point{
		X: target.X - curr.X,
		Y: target.Y - curr.Y,
	}
}

func findEmpty() (rows []int, cols []int) {
	rowlen := len(lines[0])
	hash := rune("#"[0])
	rows = []int{}
	cols = []int{}

	for c := 0; c < rowlen; c++ {
		colempty := true
		for _, row := range lines {
			if rune(row[c]) == hash {
				colempty = false
				break
			}
		}
		if colempty {
			cols = append(cols, c)
		}
	}

	for rowidx, row := range lines {
		if strings.Contains(row, "#") {
			continue
		}
		rows = append(rows, rowidx)
	}

	return rows, cols
}

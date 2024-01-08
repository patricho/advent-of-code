package main

import (
	"fmt"
	"strings"

	s "github.com/patricho/advent-of-code/go/shared"
)

type Instruction struct {
	Direction s.Point
	Len       int
}

func main() {
	testcases := map[string]int{
		"test1.txt": 6,
		"test2.txt": 159,
		"test3.txt": 135,
		"input.txt": 0,
	}

	for filename, expect := range testcases {
		func(filename string, expect int) {
			s.Measure(func() {
				got := solve(filename)
				fmt.Println(filename, "want", expect, "got", got)
			})
		}(filename, expect)
	}
}

func solve(filename string) int {
	lines := s.ReadFile(filename)
	instr1 := parseInstructions(lines[0])
	instr2 := parseInstructions(lines[1])

	wire1 := tracePositions(instr1)
	wire2 := tracePositions(instr2)

	// fmt.Println("wire1", wire1)
	// fmt.Println("wire2", wire2)

	intersections := []s.Point{}

	for _, w1 := range wire1 {
		for _, w2 := range wire2 {
			if w1 == w2 && (w1.X != 0 || w1.Y != 0) {
				intersections = append(intersections, w1)
			}
		}
	}

	// fmt.Println("intersections:", intersections)

	mindist := float64(999_999)
	origo := s.Point{X: 0, Y: 0}

	for _, intsec := range intersections {
		dist := s.ManhattanDistancePoint(origo, intsec)
		if dist < mindist {
			mindist = dist
		}
		// fmt.Println(intsec, "distance", dist)
	}

	// fmt.Println("min distance", mindist)

	return int(mindist)
}

/* func findEdges(wire1, wire2 []s.Point) (s.Point, s.Point) {
	tl, br := s.Point{}, s.Point{}

	for _, p := range append(wire1, wire2...) {
		if p.X < tl.X {
			tl.X = p.X
		}
		if p.Y < tl.Y {
			tl.Y = p.Y
		}

		if p.X > br.X {
			br.X = p.X
		}
		if p.Y > br.Y {
			br.Y = p.Y
		}
	}

	return tl, br
} */

func tracePositions(instructions []Instruction) []s.Point {
	grid := []s.Point{}
	x, y := 0, 0

	grid = append(grid, s.Point{X: x, Y: y})

	for _, instr := range instructions {
		for i := 0; i < instr.Len; i++ {
			x += instr.Direction.X
			y += instr.Direction.Y
			grid = append(grid, s.Point{X: x, Y: y})
		}
	}
	return grid
}

func parseInstructions(input string) []Instruction {
	dirs := map[string]s.Point{
		"U": s.Up,
		"D": s.Down,
		"L": s.Left,
		"R": s.Right,
	}
	strs := strings.Split(input, ",")
	out := make([]Instruction, len(strs))
	for i, str := range strs {
		dir := str[0:1]
		len := str[1:]
		out[i] = Instruction{
			Direction: dirs[dir],
			Len:       s.ToInt(len),
		}
		// fmt.Println(str, "->", dir, len, out[i])
	}
	return out
}

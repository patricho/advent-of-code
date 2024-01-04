package main

import (
	"fmt"
	"strconv"
	"strings"

	s "github.com/patricho/advent-of-code/go/shared"
)

type Instruction struct {
	Dir      string
	Len      int
	ColorDir string
	ColorLen int
}

func main() {
	s.Measure(func() { solve(1, "test.txt") })
	s.Measure(func() { solve(2, "test.txt") })

	s.Measure(func() { solve(1, "input.txt") })
	s.Measure(func() { solve(2, "input.txt") })
}

func solve(part int, filename string) {
	edges, perim := parseInput(part, filename)

	// inner area according to the shoelace formula
	ia := s.Shoelace(edges)

	// total area using pick's theorem
	area := ia + (perim / 2) + 1

	fmt.Println("part", part, filename, "total area:", area)
}

func parseInput(part int, filename string) ([]s.Point, int) {
	lines := s.ReadFile(filename)

	instructions := []Instruction{}

	for _, line := range lines {
		parts := strings.Split(line, " ")

		// (#70c710)
		color := parts[2]
		colorDir := color[7:8] // inclusive to the left, exclusive to the right
		colorLen, _ := strconv.ParseInt(color[2:7], 16, 32)

		instructions = append(instructions, Instruction{
			Dir:      parts[0],
			Len:      s.ToInt(parts[1]),
			ColorDir: colorDir,
			ColorLen: int(colorLen),
		})
	}

	dirs := map[string]s.Point{
		"R": s.Right,
		"L": s.Left,
		"U": s.Up,
		"D": s.Down,

		"0": s.Right,
		"2": s.Left,
		"3": s.Up,
		"1": s.Down,
	}

	edges := []s.Point{
		{X: 0, Y: 0},
	}

	prev := s.Point{X: 0, Y: 0}
	perim := 0

	for _, i := range instructions {
		var dir s.Point
		var len int

		if part == 1 {
			dir = dirs[i.Dir]
			len = i.Len
		} else {
			dir = dirs[i.ColorDir]
			len = i.ColorLen
		}

		e := s.Point{
			X: prev.X + dir.X*len,
			Y: prev.Y + dir.Y*len,
		}

		edges = append(edges, e)
		perim += len
		prev = e
	}

	return edges, perim
}

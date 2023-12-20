package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/patricho/advent-of-code/go/shared"
)

/*
| is a vertical pipe connecting north and south: │
- is a horizontal pipe connecting east and west: ─
L is a 90-degree bend connecting north and east: └
J is a 90-degree bend connecting north and west: ┘
7 is a 90-degree bend connecting south and west: ┐
F is a 90-degree bend connecting south and east: ┌
*/

type Point struct {
	Char string
	X    int
	Y    int
}

var (
	dirs  map[string][]Point
	lines []string
)

func main() {
	lines = shared.ReadFile("input.txt")
	up := Point{X: 0, Y: -1}
	left := Point{X: -1, Y: 0}
	right := Point{X: 1, Y: 0}
	down := Point{X: 0, Y: 1}
	dirs = map[string][]Point{
		"S": {right},
		"│": {up, down},
		"|": {up, down},
		"─": {left, right},
		"-": {left, right},
		"└": {up, right},
		"L": {up, right},
		"┘": {left, up},
		"J": {left, up},
		"┐": {left, down},
		"7": {left, down},
		"┌": {right, down},
		"F": {right, down},
	}
	// prettifyInput(lines)
	start := findStart()
	// fmt.Println("start", start)
	visited := []Point{}
	enclosed := []Point{}
	visit(start, start, start, &visited)
	printLines(prettifyVisited(&visited, true))
	fmt.Println("steps total:", len(visited), "furthest:", len(visited)/2)

	// part 2
	lines = prettifyVisited(&visited, false)
	calculateEnclosed(visited, &enclosed)
}

func printLines(prlines []string) {
	for _, line := range prlines {
		fmt.Println(line)
	}
}

func calculateEnclosed(visited []Point, enclosed *[]Point) {
	n := 0

	cg := color.New(color.FgGreen)
	ce := color.New(color.FgRed) //.Add(color.Faint)
	cf := color.New(color.Faint)

	for _, line := range lines {
		passages := 0
		inF, inL := false, false
		for _, char := range line {
			printed := false
			switch string(char) {
			case "|":
				// definitely a passage
				passages++
			case "-":
				// neutral
			case "F":
				// passage in potentia
				inF = true
			case "L", "S": // HACK: manually checked input for what S is...
				// passage in potentia
				inL = true
			case "7":
				if inL {
					// L7 now were passing
					inL = false
					passages++
				}
				inF = false
			case "J":
				if inF {
					// FJ now were passing
					inF = false
					passages++
				}
				inL = false
			default:
				inside := passages%2 != 0
				if inside {
					n++
					ce.Print("O")
				} else {
					cf.Print(".")
				}
				printed = true
			}

			if !printed {
				cg.Print(prettifyChar(char))
			}
		}
		fmt.Print("\n")
	}
	fmt.Println("enclosed", n)
}

func visit(start, curr, prev Point, visited *[]Point) {
	// fmt.Println("visit, curr", curr, "dirs", dirs[curr.Char])

	if curr.X == start.X && curr.Y == start.Y && len(*visited) > 0 {
		// fmt.Println("back at start!")
		return
	}

	*visited = append(*visited, curr)

	next := Point{X: -1, Y: -1}
	for _, dir := range dirs[curr.Char] {
		poss := Point{X: curr.X + dir.X, Y: curr.Y + dir.Y}
		poss.Char = string(rune(lines[poss.Y][poss.X]))

		// fmt.Println("curr", curr, "dir", dir, "poss", poss)

		if poss.X != prev.X || poss.Y != prev.Y {
			// fmt.Println("found next!", poss)
			next = poss
			break
		}
	}
	if next.X < 0 || next.Y < 0 {
		panic("wtf")
	}

	// fmt.Println("next", next)

	visit(start, next, curr, visited)
}

func isVisited(visited *[]Point, x, y int) int {
	for idx, node := range *visited {
		if node.X == x && node.Y == y {
			return idx
		}
	}
	return -1
}

func prettifyVisited(visited *[]Point, prettyPath bool) []string {
	newlines := []string{}

	for y, line := range lines {
		newline := ""
		for x, char := range line {
			if visidx := isVisited(visited, x, y); visidx >= 0 {
				if prettyPath {
					newline += prettifyChar(char)
				} else {
					newline += string(char)
				}
			} else {
				newline += "."
			}
		}
		newlines = append(newlines, newline)
	}

	return newlines
}

func findStart() Point {
	for y, line := range lines {
		for x, char := range line {
			if string(char) == "S" {
				return Point{Char: "S", X: x, Y: y}
			}
		}
	}
	panic("wtf")
}

func prettifyChar(in rune) string {
	if string(in) == "|" {
		return "│"
	}
	if string(in) == "-" {
		return "─"
	}
	if string(in) == "L" {
		return "└"
	}
	if string(in) == "J" {
		return "┘"
	}
	if string(in) == "7" {
		return "┐"
	}
	if string(in) == "F" {
		return "┌"
	}
	if string(in) == "S" {
		return "S"
	}
	return "."
}

func prettifyInput(lines []string) {
	for _, line := range lines {
		newline := ""
		for _, char := range line {
			newline += prettifyChar(char)
		}
		fmt.Println(newline)
	}
}

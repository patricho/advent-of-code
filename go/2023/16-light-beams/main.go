package main

import (
	"fmt"

	"github.com/fatih/color"
	s "github.com/patricho/advent-of-code/go/shared"
)

type Beam struct {
	s.Point
	Direction s.Point
}

var grid [][]rune

func main() {
	lines := s.ReadFile("input.txt")
	grid = s.LinesToGrid(lines)
	beams := map[Beam]bool{}
	visited := map[s.Point]bool{}
	start := Beam{
		Point:     s.Point{X: 0, Y: 0},
		Direction: s.Right,
	}
	queue := s.CreateQueue[Beam]()
	queue.Enqueue(start)

	rows, cols := len(grid), len(grid[0])

	for queue.Len() > 0 {
		current := queue.Dequeue()
		beams[current] = true
		visited[current.Point] = true

		currChar := grid[current.Y][current.X]

		dirs := []s.Point{}
		dirsstr := ""

		if currChar == '|' && (current.Direction == s.Right || current.Direction == s.Left) {
			dirs = append(dirs, s.Up, s.Down)
			dirsstr += "up/down"
		} else if currChar == '-' && (current.Direction == s.Down || current.Direction == s.Up) {
			dirs = append(dirs, s.Left, s.Right)
			dirsstr += "left/right"
		} else if currChar == '/' {
			if current.Direction == s.Down {
				dirs = append(dirs, s.Left)
				dirsstr += "left"
			} else if current.Direction == s.Up {
				dirs = append(dirs, s.Right)
				dirsstr += "right"
			} else if current.Direction == s.Left {
				dirs = append(dirs, s.Down)
				dirsstr += "down"
			} else if current.Direction == s.Right {
				dirs = append(dirs, s.Up)
				dirsstr += "up"
			}
		} else if currChar == '\\' {
			if current.Direction == s.Down {
				dirs = append(dirs, s.Right)
				dirsstr += "right"
			} else if current.Direction == s.Up {
				dirs = append(dirs, s.Left)
				dirsstr += "left"
			} else if current.Direction == s.Left {
				dirs = append(dirs, s.Up)
				dirsstr += "up"
			} else if current.Direction == s.Right {
				dirs = append(dirs, s.Down)
				dirsstr += "down"
			}
		} else {
			dirs = append(dirs, current.Direction)
			dirsstr += "unchanged"
		}

		// fmt.Println(current.Point, string(currChar), dirsstr, dirs)
		// printGrid(visited, current.Point)
		// fmt.Println("")

		for _, dir := range dirs {
			next := Beam{
				Point: s.Point{
					X: current.X + dir.X,
					Y: current.Y + dir.Y,
				},
				Direction: dir,
			}

			if next.Y < 0 || next.Y >= rows || next.X < 0 || next.X >= cols {
				// Out of bounds
				continue
			} else if beams[next] {
				// Already visited (in this direction)
				continue
			}

			queue.Enqueue(next)
		}
	}

	printGrid(visited, s.Point{X: -1, Y: -1})

	fmt.Println("energized:", len(visited))
}

func printGrid(visited map[s.Point]bool, hl s.Point) {
	gray := color.New(color.FgBlue).Add(color.Faint)
	red := color.New(color.FgRed).Add(color.Bold)
	yellow := color.New(color.FgYellow)

	for y, row := range grid {
		for x, c := range row {
			vis := visited[s.Point{X: x, Y: y}]
			if y == hl.Y && x == hl.X {
				yellow.Print(string(c))
			} else if vis {
				red.Print(string(c))
			} else {
				gray.Print(string(c))
			}
		}
		fmt.Print("\n")
	}
}

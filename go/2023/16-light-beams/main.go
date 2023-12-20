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
	s.Measure(part1)
	s.Measure(part2)
}

func part1() {
	start := Beam{
		Point:     s.Point{X: -1, Y: 0},
		Direction: s.Right,
	}
	visited := solve(start)
	printGrid(visited)
	fmt.Println("part 1 energized:", len(visited))
}

func part2() {
	rows, cols := len(grid), len(grid[0])
	maxEnergized := 0
	starts := []Beam{}

	for row := 0; row < rows; row++ {
		starts = append(
			starts,
			Beam{ // left to right
				Point:     s.Point{X: -1, Y: row},
				Direction: s.Right,
			},
			Beam{ // right to left
				Point:     s.Point{X: cols, Y: row},
				Direction: s.Left,
			},
		)
	}

	for col := 0; col < cols; col++ {
		starts = append(
			starts,
			Beam{ // top to bottom
				Point:     s.Point{X: col, Y: -1},
				Direction: s.Down,
			},
			Beam{ // bottom to top
				Point:     s.Point{X: col, Y: rows},
				Direction: s.Up,
			},
		)
	}

	for _, start := range starts {
		visited := solve(start)
		if len(visited) > maxEnergized {
			maxEnergized = len(visited)
		}
	}

	fmt.Println("part 2 max energized", maxEnergized)
}

func solve(start Beam) map[s.Point]bool {
	beams := map[Beam]bool{}
	visited := map[s.Point]bool{}
	queue := s.CreateQueue[Beam]()
	queue.Enqueue(start)

	for queue.Len() > 0 {
		current := queue.Dequeue()

		currChar := '.'
		if !s.OOB(grid, current.Point) {
			beams[current] = true
			visited[current.Point] = true
			currChar = grid[current.Y][current.X]
		}

		dirs := []s.Point{}

		if currChar == '|' && (current.Direction == s.Right || current.Direction == s.Left) {
			dirs = append(dirs, s.Up, s.Down)
		} else if currChar == '-' && (current.Direction == s.Down || current.Direction == s.Up) {
			dirs = append(dirs, s.Left, s.Right)
		} else if currChar == '/' {
			switch current.Direction {
			case s.Down:
				dirs = append(dirs, s.Left)
			case s.Up:
				dirs = append(dirs, s.Right)
			case s.Left:
				dirs = append(dirs, s.Down)
			case s.Right:
				dirs = append(dirs, s.Up)
			}
		} else if currChar == '\\' {
			switch current.Direction {
			case s.Down:
				dirs = append(dirs, s.Right)
			case s.Up:
				dirs = append(dirs, s.Left)
			case s.Left:
				dirs = append(dirs, s.Up)
			case s.Right:
				dirs = append(dirs, s.Down)
			}
		} else {
			dirs = append(dirs, current.Direction)
		}

		for _, dir := range dirs {
			next := Beam{
				Point: s.Point{
					Y: current.Y + dir.Y,
					X: current.X + dir.X,
				},
				Direction: dir,
			}

			if s.OOB(grid, next.Point) {
				continue
			} else if beams[next] {
				// already visited (in this direction)
				continue
			}

			queue.Enqueue(next)
		}
	}

	return visited
}

func printGrid(visited map[s.Point]bool) {
	gray := color.New(color.FgBlue).Add(color.Faint)
	red := color.New(color.FgRed).Add(color.Bold)

	for y, row := range grid {
		for x, c := range row {
			vis := visited[s.Point{X: x, Y: y}]
			/*if y == hl.Y && x == hl.X {
				yellow.Print(string(c))
			} else*/
			if vis {
				red.Print(string(c))
			} else {
				gray.Print(string(c))
			}
		}
		fmt.Print("\n")
	}
}

package main

import (
	"fmt"

	"github.com/fatih/color"
	s "github.com/patricho/advent-of-code/go/shared"
)

func main() {
	lines := s.ReadFile("input.txt")

	grid := [][]rune{}
	for _, line := range lines {
		row := []rune{}
		for _, c := range line {
			row = append(row, c)
		}
		grid = append(grid, row)
	}

	start := findPoint(grid, 'S')
	end := findPoint(grid, 'E')

	path := findPath(grid, start, end)

	if path != nil {
		fmt.Println("path length:", len(path), ", steps:", len(path)-1)

		gray := color.New(color.FgBlue).Add(color.Faint)
		green := color.New(color.FgRed).Add(color.Bold)

		for y, row := range grid {
			for x, c := range row {
				visited := false
				for _, p := range path {
					if p.X == x && p.Y == y {
						visited = true
						break
					}
				}
				if visited {
					green.Print(string(c))
				} else {
					gray.Print(string(c))
				}
			}
			fmt.Print("\n")
		}
	} else {
		fmt.Println("no path found")
	}
}

func findPoint(grid [][]rune, r rune) s.Point {
	for y, row := range grid {
		for x, c := range row {
			if c == r {
				return s.Point{
					X: x,
					Y: y,
				}
			}
		}
	}
	panic("wtf")
}

// findPath is a modified BFS algorithm with specific rules for this use case
func findPath(grid [][]rune, start, end s.Point) []s.Point {
	rows, cols := len(grid), len(grid[0])

	queue := s.CreateQueue[s.Point]()
	visited := make(map[s.Point]bool)
	parent := make(map[s.Point]s.Point)

	queue.Enqueue(start)
	visited[start] = true

	for queue.Len() > 0 {
		current := queue.Dequeue()

		if current == end {
			// Reconstruct the path from end to start
			var path []s.Point
			for current != start {
				path = append([]s.Point{current}, path...)
				current = parent[current]
			}
			path = append([]s.Point{start}, path...)
			return path
		}

		currentChar := grid[current.Y][current.X]

		for _, move := range s.Directions {
			next := s.Point{X: current.X + move.X, Y: current.Y + move.Y}

			if next.Y < 0 || next.Y >= rows || next.X < 0 || next.X >= cols {
				// Out of bounds
				continue
			} else if visited[next] {
				// Already visited
				continue
			}

			nextChar := grid[next.Y][next.X]

			// For the ASCII diff calculation to be correct
			if currentChar == 'S' {
				currentChar = 'a'
			} else if nextChar == 'E' {
				nextChar = 'z'
			}

			charDiff := nextChar - currentChar

			if charDiff > 1 {
				// Too big a leap
				continue
			}

			queue.Enqueue(next)

			visited[next] = true
			parent[next] = current
		}
	}

	return nil
}

package main

import (
	"fmt"

	"github.com/fatih/color"
	s "github.com/patricho/advent-of-code/go/shared"
)

var grid [][]rune

func main() {
	parseGrid()
	// part1()
	part2()
	// test()
}

func parseGrid() {
	lines := s.ReadFile("input.txt")
	grid = [][]rune{}
	for _, line := range lines {
		row := []rune{}
		for _, c := range line {
			row = append(row, c)
		}
		grid = append(grid, row)
	}
}

func part2() {
	starts := findPoints('a')
	end := findPoints('E')[0]
	shortest := 999999
	shortestPath := []s.Point{}
	for _, start := range starts {
		path := findPath(start, end)
		if len(path) > 0 && len(path) < shortest {
			// printPath(path)
			shortest = len(path)
			shortestPath = path
		}
	}
	printPath(shortestPath)
	fmt.Println("shortest path found:", shortest, "steps:", shortest-1)
}

func part1() {
	start := findPoints('S')[0]
	end := findPoints('E')[0]

	path := findPath(start, end)

	if path == nil {
		panic("no path found")
	}

	printPath(path)

	fmt.Println("path length:", len(path), ", steps:", len(path)-1)
}

func printPath(path []s.Point) {
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
}

func findPoints(r rune) []s.Point {
	output := []s.Point{}
	for y, row := range grid {
		for x, c := range row {
			if c == r {
				output = append(output, s.Point{
					X: x,
					Y: y,
				})
			}
		}
	}
	return output
}

// findPath is a modified BFS algorithm with specific rules for this use case
func findPath(start, end s.Point) []s.Point {
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

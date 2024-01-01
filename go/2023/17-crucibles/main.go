package main

import (
	"container/heap"
	"fmt"
	"slices"

	"github.com/fatih/color"
	s "github.com/patricho/advent-of-code/go/shared"
)

type DirectedPoint struct {
	X, Y, DirX, DirY, Streak int
}

type Node struct {
	s.Point
	Cost      int
	Direction s.Point
	Streak    int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Cost < pq[j].Cost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

var (
	grid       [][]int
	rows, cols int
	start, end s.Point
)

func main() {
	s.Measure(func() { solve(1, "test1.txt") })
	s.Measure(func() { solve(2, "test1.txt") })
	s.Measure(func() { solve(2, "test2.txt") })
	s.Measure(func() { solve(1, "input.txt") })
	s.Measure(func() { solve(2, "input.txt") })
	s.Measure(func() { solve(3, "input.txt") })
}

func solve(part int, filename string) {
	lines := s.ReadFile(filename)
	grid = s.LinesToNumberGrid(lines)
	rows, cols = len(grid), len(grid[0])

	start = s.Point{Y: 0, X: 0}
	end = s.Point{Y: rows - 1, X: cols - 1}

	minStreak := 0
	maxStreak := 3
	if part == 2 {
		minStreak = 4
		maxStreak = 10
	} else if part == 3 {
		minStreak = 0
		maxStreak = 999
	}

	path, cost := findPath(grid, start, end, minStreak, maxStreak)
	if path != nil {
		fmt.Println("part", part, filename, "cost", cost)
	} else {
		fmt.Println("part", part, filename, "no path found")
	}
}

func findPath(grid [][]int, start, end s.Point, minStreak, maxStreak int) ([]s.Point, int) {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	visited := make(map[DirectedPoint]bool)
	costs := make(map[DirectedPoint]int)
	parent := make(map[DirectedPoint]DirectedPoint)
	steps := 0

	// Enqueue the starting point with cost 0 (twice, once in each possible direction)
	heap.Push(&pq, &Node{Point: start, Cost: 0, Direction: s.Point{X: 1, Y: 0}, Streak: 0})
	startd1 := DirectedPoint{X: start.X, Y: start.Y, DirX: 1, DirY: 0, Streak: 0}
	costs[startd1] = 0
	visited[startd1] = true

	heap.Push(&pq, &Node{Point: start, Cost: 0, Direction: s.Point{X: 0, Y: 1}, Streak: 0})
	startd2 := DirectedPoint{X: start.X, Y: start.Y, DirX: 0, DirY: 1, Streak: 0}
	costs[startd2] = 0
	visited[startd2] = true

	for pq.Len() > 0 {
		currentNode := heap.Pop(&pq).(*Node)
		current := currentNode.Point
		currentd := DirectedPoint{X: current.X, Y: current.Y, DirX: currentNode.Direction.X, DirY: currentNode.Direction.Y, Streak: currentNode.Streak}
		steps++

		if current == end {
			fmt.Println("total cells:", cols*rows, ", steps taken:", steps, ", parents:", len(parent), ", cost:", costs[currentd])
			endd := currentd
			cost := costs[currentd]

			// Reconstruct the path from end to start
			var path []s.Point
			for current != start {
				path = append([]s.Point{current}, path...)
				currentd = parent[currentd]
				current = s.Point{X: currentd.X, Y: currentd.Y}
			}
			path = append([]s.Point{start}, path...)

			// Show the path taken for fun
			showGrid(endd, path)

			return path, cost
		}

		for _, move := range s.Directions {
			next := s.Point{X: current.X + move.X, Y: current.Y + move.Y}
			nextd := DirectedPoint{X: next.X, Y: next.Y, DirX: move.X, DirY: move.Y, Streak: 0}

			if s.OOB(grid, next) {
				continue
			}

			nextd.Streak = 0

			if next == end && currentNode.Streak+1 < minStreak-1 {
				// End found, but with too short a streak before
				continue
			} else if move == currentNode.Direction {
				if currentNode.Streak >= maxStreak-1 {
					// Can't continue straight on
					continue
				}
				nextd.Streak = currentNode.Streak + 1
			} else if move.X == currentNode.Direction.X && move.Y != currentNode.Direction.Y {
				// reversing
				continue
			} else if move.Y == currentNode.Direction.Y && move.X != currentNode.Direction.X {
				// reversing
				continue
			} else if currentd.Streak < minStreak-1 {
				// Can't turn yet
				continue
			}

			// Calculate the cost to reach the next point
			nextCost := costs[currentd] + grid[next.Y][next.X]

			// Check if the cost to reach the next point is less than the current known cost
			if !visited[nextd] || nextCost < costs[nextd] {
				// Enqueue the next point with the updated cost
				heap.Push(&pq, &Node{next, nextCost, move, nextd.Streak})
				costs[nextd] = nextCost
				visited[nextd] = true
				parent[nextd] = currentd
			}
		}
	}

	// No path found
	return nil, 0
}

func showGrid(current DirectedPoint, path []s.Point) {
	blue := color.New(color.FgBlue).Add(color.Faint)
	red := color.New(color.FgRed).Add(color.Bold)
	yellow := color.New(color.FgYellow).Add(color.Bold)

	currentpt := s.Point{X: current.X, Y: current.Y}

	for y, row := range grid {
		for x, c := range row {
			pt := s.Point{Y: y, X: x}
			inpath := slices.Contains(path, pt)

			if pt == currentpt {
				yellow.Print(fmt.Sprint(c))
			} else if inpath {
				red.Print(fmt.Sprint(c))
			} else {
				blue.Print(fmt.Sprint(c))
			}
		}
		fmt.Print("\n")
	}
}

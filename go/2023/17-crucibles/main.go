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
	start, end s.Point
)

func main() {
	lines := s.ReadFile("test.txt")
	grid = s.LinesToNumberGrid(lines)

	start = s.Point{Y: 0, X: 0}
	end = s.Point{Y: 12, X: 12}

	path, cost := findPath(grid, start, end)
	if path != nil {
		fmt.Printf("Path found with cost %d:\n", cost)
		for _, point := range path {
			fmt.Printf("(%d, %d) ", point.X, point.Y)
		}
		fmt.Println("")

	} else {
		fmt.Println("No path found.")
	}
}

func findPath(grid [][]int, start, end s.Point) ([]s.Point, int) {
	rows, cols := len(grid), len(grid[0])

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	visited := make(map[DirectedPoint]bool)
	costs := make(map[DirectedPoint]int)
	parent := make(map[DirectedPoint]DirectedPoint)
	steps := 0

	// Enqueue the starting point with cost 0
	heap.Push(&pq, &Node{Point: start, Cost: 0, Direction: s.Point{X: 1, Y: 0}, Streak: 0})
	startd1 := DirectedPoint{X: start.X, Y: start.Y, DirX: 1, DirY: 0, Streak: 0}
	costs[startd1] = 0
	visited[startd1] = true

	heap.Push(&pq, &Node{Point: start, Cost: 0, Direction: s.Point{X: 0, Y: 1}, Streak: 0})
	startd2 := DirectedPoint{X: start.X, Y: start.Y, DirX: 0, DirY: 1, Streak: 0}
	costs[startd2] = 0
	visited[startd2] = true

	// Define possible moves (up, down, left, right)
	moves := []s.Point{
		{Y: -1, X: 0}, // up
		{Y: 1, X: 0},  // down
		{Y: 0, X: -1}, // left
		{Y: 0, X: 1},  // right
	}

	for pq.Len() > 0 {
		currentNode := heap.Pop(&pq).(*Node)
		current := currentNode.Point
		currentd := DirectedPoint{X: current.X, Y: current.Y, DirX: currentNode.Direction.X, DirY: currentNode.Direction.Y, Streak: currentNode.Streak}
		steps++

		fmt.Println("current", current.Y, current.X, grid[current.Y][current.X], "cost", costs[currentd])
		// showGrid(current, parent, visited)

		if current == end {
			fmt.Println("total cells:", cols*rows, "steps taken:", steps, "parents:", len(parent), "cost", costs[currentd])
			endd := currentd
			cost := costs[currentd]
			fmt.Println("endd:", endd)
			// Reconstruct the path from end to start
			var path []s.Point
			for current != start {
				path = append([]s.Point{current}, path...)
				currentd = parent[currentd]
				current = s.Point{X: currentd.X, Y: currentd.Y}
			}
			path = append([]s.Point{start}, path...)
			showGrid(endd, path)
			fmt.Println("endd:", endd)
			fmt.Println("returning cost", cost)
			return path, cost
		}

		// Explore possible moves
		for _, move := range moves {
			next := s.Point{X: current.X + move.X, Y: current.Y + move.Y}
			nextd := DirectedPoint{X: next.X, Y: next.Y, DirX: move.X, DirY: move.Y, Streak: 0}

			// Check if the next point is within bounds
			if next.Y < 0 || next.Y >= rows || next.X < 0 || next.X >= cols {
				fmt.Println("next", nextd, "skip - oob")
				continue
			}

			nextd.Streak = 0
			if move == currentNode.Direction {
				if currentNode.Streak >= 2 {
					// Can't continue straight on
					fmt.Println("next", nextd, grid[next.Y][next.X], "skip - needs to turn")
					continue
				}

				nextd.Streak = currentNode.Streak + 1
			} else if move.X == currentNode.Direction.X && move.Y != currentNode.Direction.Y {
				// reversing
				fmt.Println("next", nextd, grid[next.Y][next.X], "skip - reversing")
				continue
			} else if move.Y == currentNode.Direction.Y && move.X != currentNode.Direction.X {
				// reversing
				fmt.Println("next", nextd, grid[next.Y][next.X], "skip - reversing")
				continue
			}

			// Calculate the cost to reach the next point
			nextCost := costs[currentd] + grid[next.Y][next.X]

			if next == end {
				fmt.Println("one path to end found", next, nextCost)
			}

			// Check if the cost to reach the next point is less than the current known cost
			if !visited[nextd] || nextCost < costs[nextd] {
				// Enqueue the next point with the updated cost
				heap.Push(&pq, &Node{next, nextCost, move, nextd.Streak})
				costs[nextd] = nextCost
				visited[nextd] = true
				parent[nextd] = currentd
			} else {
				fmt.Println("next", nextd, grid[next.Y][next.X], "skip - cost", nextCost)
			}
		}
	}

	// No path found
	return nil, 0
}

func showGrid(current DirectedPoint, path []s.Point) {
	// gray := color.New(color.FgHiWhite).Add(color.Faint)
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
				// } else if visited[p] {
				// 	gray.Print(fmt.Sprint(c))
			} else {
				blue.Print(fmt.Sprint(c))
			}
		}
		fmt.Print("\n")
	}
}

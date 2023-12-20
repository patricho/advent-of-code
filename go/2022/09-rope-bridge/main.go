package ropebridge09

import (
	"fmt"
	"strconv"
	"strings"

	c "github.com/patricho/advent-of-code/go/util"
)

type Coordinate struct {
	x     int
	y     int
	label string
}

func Part1() {
	head := Coordinate{x: 0, y: 4}
	tail := Coordinate{x: 0, y: 4}
	visited := []Coordinate{}

	moves := c.ReadFile("09-rope-bridge/input.txt")

	for _, move := range moves {
		fmt.Println(move)

		movement, steps := parseMove(move)

		for i := 1; i <= steps; i++ {
			moveHead(&head, movement)
			moveTail(head, &tail)
			// draw(head, tail)

			visited = append(visited, Coordinate{x: tail.x, y: tail.y})
		}
	}

	fmt.Println("\n", "final result", "\n")

	drawVisited(visited)
}

func Part2() {
	parts := []Coordinate{
		{x: 11, y: 15, label: "H"},
		{x: 11, y: 15, label: "1"},
		{x: 11, y: 15, label: "2"},
		{x: 11, y: 15, label: "3"},
		{x: 11, y: 15, label: "4"},
		{x: 11, y: 15, label: "5"},
		{x: 11, y: 15, label: "6"},
		{x: 11, y: 15, label: "7"},
		{x: 11, y: 15, label: "8"},
		{x: 11, y: 15, label: "9"},
	}

	visited := []Coordinate{}

	moves := c.ReadFile("09-rope-bridge/input.txt")

	for _, move := range moves {
		fmt.Println("")
		fmt.Println(move)

		movement, steps := parseMove(move)

		for i := 1; i <= steps; i++ {
			// fmt.Println("")

			moveHead(&parts[0], movement)

			for j := 1; j < len(parts); j++ {
				moveTail(parts[j-1], &parts[j])
			}

			// drawAll(parts)

			visited = append(visited, Coordinate{x: parts[9].x, y: parts[9].y})
		}

		// drawAll(parts)
	}

	fmt.Println("\n", "final result", "\n")

	drawVisited(visited)
}

func drawVisited(visited []Coordinate) {
	// find coordinates range
	xmin, xmax, ymin, ymax := 0, 0, 0, 0

	for _, v := range visited {
		if v.x < xmin {
			xmin = v.x
		}
		if v.x > xmax {
			xmax = v.x
		}
		if v.y < ymin {
			ymin = v.y
		}
		if v.y > ymax {
			ymax = v.y
		}
	}

	visitedCount := 0

	for r := ymin; r <= ymax; r++ {
		for c := xmin; c <= xmax; c++ {
			found := false

			for _, v := range visited {
				if v.x == c && v.y == r {
					found = true
					break
				}
			}

			if found {
				fmt.Print("#")
				visitedCount++
			} else {
				fmt.Print(".")
			}

		}

		fmt.Print("\n")
	}

	fmt.Println("visitedCount", visitedCount)
}

func parseMove(move string) (Coordinate, int) {
	parts := strings.Split(move, " ")

	var movement Coordinate

	if parts[0] == "R" {
		movement = Coordinate{x: 1, y: 0}
	} else if parts[0] == "L" {
		movement = Coordinate{x: -1, y: 0}
	} else if parts[0] == "U" {
		movement = Coordinate{x: 0, y: -1}
	} else if parts[0] == "D" {
		movement = Coordinate{x: 0, y: 1}
	}

	steps, _ := strconv.Atoi(parts[1])

	return movement, steps
}

func draw(head, tail Coordinate) {
	rows, cols := 5, 6

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if head.x == c && head.y == r {
				fmt.Print("H")
			} else if tail.x == c && tail.y == r {
				fmt.Print("T")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Print("\n")
	}
}

func drawAll(parts []Coordinate) {
	rows, cols := 21, 26

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			found := false

			for _, part := range parts {
				if part.x == c && part.y == r && !found {
					fmt.Print(part.label)
					found = true
				}
			}

			if !found {
				fmt.Print(".")
			}

		}

		fmt.Print("\n")
	}
}

func moveTail(head Coordinate, tail *Coordinate) {
	delta := Coordinate{x: head.x - (*tail).x, y: head.y - (*tail).y}
	movement := Coordinate{x: 0, y: 0}

	if c.Abs(delta.x) == 2 {
		// if x diff is 2 steps, move 1 to keep up
		movement.x = delta.x / 2

		// if there's also an y diff, move 1 to keep up
		if delta.y >= 1 {
			movement.y = 1
		} else if delta.y <= -1 {
			movement.y = -1
		}
	} else if c.Abs(delta.y) == 2 {
		// if y diff is 2 steps, move 1 to keep up
		movement.y = delta.y / 2

		// if there's also an x diff, move 1 to keep up
		if delta.x >= 1 {
			movement.x = 1
		} else if delta.x <= -1 {
			movement.x = -1
		}
	}

	/*if (*tail).label == "5" {
		fmt.Println("head", head)
		fmt.Println("tail", *tail)
		fmt.Println("delta", delta)
		fmt.Println("movement", movement)
	}*/

	(*tail).x += movement.x
	(*tail).y += movement.y

	/*if (*tail).label == "5" {
		fmt.Println("result", "head", head, "tail", *tail)
	}*/
}

func moveHead(head *Coordinate, movement Coordinate) {
	(*head).x += movement.x
	(*head).y += movement.y
}

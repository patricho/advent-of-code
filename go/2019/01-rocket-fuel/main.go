package main

import (
	"fmt"
	"math"

	"github.com/patricho/advent-of-code/go/shared"
)

func main() {
	solve(2, "test.txt")
	solve(2, "input.txt")
}

func solve(part int, filename string) {
	lines := shared.ReadFile(filename)
	sum := int64(0)
	for _, line := range lines {
		sum += processLine(part, line)
	}
	fmt.Println(part, filename, "sum:", sum)
}

func processLine(part int, line string) int64 {
	n := calc(shared.ToInt64(line))

	// fmt.Println(n)

	if part == 2 {
		nn := n
		for nn > 0 {
			nn = calc(nn)
			// fmt.Println("  ", nn)
			if nn > 0 {
				n += nn
			}
		}
		// fmt.Println(n)
		// fmt.Println("")
	}

	return int64(n)
}

func calc(input int64) int64 {
	return int64(math.Floor(float64(input)/3) - 2)
}

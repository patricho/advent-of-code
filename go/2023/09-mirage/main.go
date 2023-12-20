package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/patricho/advent-of-code/go/shared"
)

func main() {
	run(1, "input.txt")
	run(2, "input.txt")
}

func run(part int, filename string) {
	lines := shared.ReadFile(filename)

	histories := [][][]int{}
	for _, line := range lines {
		nums := []int{}
		for _, n := range strings.Split(line, " ") {
			nums = append(nums, shared.ToInt(n))
		}
		if part == 2 {
			slices.Reverse(nums)
		}
		histories = append(histories, [][]int{nums})
	}

	added := 0

	for _, history := range histories {
		diff := true

		for i := 0; i < len(history) && diff; i++ {
			row := history[i]
			diffs := []int{}
			for j := 1; j < len(row); j++ {
				diff := row[j] - row[j-1]
				diffs = append(diffs, diff)
			}
			diff = !shared.All(row, 0)
			if diff {
				history = append(history, diffs)
			}
		}

		for i := len(history) - 1; i >= 0; i-- {
			rowlen := len(history[i])

			// last row - add 0
			if i == len(history)-1 {
				// fmt.Println(iter, "row", history[i], "add 0")
				history[i] = append(history[i], 0)
				continue
			}

			// add from the row below
			prevrow := history[i+1]
			add := shared.Last(prevrow) + shared.Last(history[i])
			// fmt.Println(iter, "row", history[i], "add", last(prevrow), "+", last(history[i]), "=", add)
			history[i] = append(history[i], add)
			rowlen++

		}

		added += shared.Last(history[0])

		// fmt.Println(iter, "row", last(history[0]))
	}

	fmt.Println("part", part, "added total sum:", added)
}

package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/patricho/advent-of-code/go/shared"
)

func main() {
	shared.DisplayMain(func() {
		test()
		run()
	})
}

func test() {
	shared.RunCase("part 1 test", func() int {
		return part1("../inputs/2025/01-test.txt")
	}, 3)

	shared.RunCase("part 2 test", func() int {
		return part2("../inputs/2025/01-test.txt")
	}, 6)

	// // A bunch of test cases for part 2
	// shared.RunCase("part 2 test case", func() int { return part2Loop([]int{1000}) }, 10)
	// //  All these should give exactly 1
	// shared.RunCase("part 2 test case", func() int { return part2Loop([]int{-75, 20}) }, 1)
	// shared.RunCase("part 2 test case", func() int { return part2Loop([]int{75, -20}) }, 1)
	// shared.RunCase("part 2 test case", func() int { return part2Loop([]int{-50, 50}) }, 1)
	// shared.RunCase("part 2 test case", func() int { return part2Loop([]int{-50, -50}) }, 1)
	// shared.RunCase("part 2 test case", func() int { return part2Loop([]int{50, 50}) }, 1)
	// shared.RunCase("part 2 test case", func() int { return part2Loop([]int{50, -50}) }, 1)
	// //  All these should give exactly 2
	// shared.RunCase("part 2 test case", func() int { return part2Loop([]int{-200}) }, 2)
	// shared.RunCase("part 2 test case", func() int { return part2Loop([]int{200}) }, 2)
	// shared.RunCase("part 2 test case", func() int { return part2Loop([]int{-150, -50}) }, 2)
	// shared.RunCase("part 2 test case", func() int { return part2Loop([]int{-150, 50}) }, 2)
	// shared.RunCase("part 2 test case", func() int { return part2Loop([]int{150, -50}) }, 2)
	// shared.RunCase("part 2 test case", func() int { return part2Loop([]int{150, 50}) }, 2)
	// shared.RunCase("part 2 test case", func() int { return part2Loop([]int{-50, 100}) }, 2)
}

func run() {
	shared.RunCase("part 1", func() int {
		return part1("../inputs/2025/01-input.txt")
	}, 1092)

	shared.RunCase("part 2", func() int {
		return part2("../inputs/2025/01-input.txt")
	}, 6616)
}

func part1(filename string) int {
	zeroes := 0
	position := 50

	for _, step := range getSteps(filename) {
		position += step

		if position%100 == 0 {
			zeroes++
		}
	}

	return zeroes
}

func part2(filename string) int {
	return part2Loop(getSteps(filename))
}

func part2Loop(steps []int) int {
	zeroes := 0
	position := 50

	for _, step := range steps {
		oldposition := position

		position += step

		oldround := int(math.Floor(float64(oldposition) / float64(100)))
		posround := int(math.Floor(float64(position) / float64(100)))

		diff := int(math.Abs(float64(posround - oldround)))

		if step < 0 && position%100 == 0 {
			// Stepping backwards, landing on 0 - count that
			diff++
		} else if step < 0 && oldposition%100 == 0 {
			// Stepping backwards, starting at 0 - already counted that
			diff--
		}

		zeroes += diff
	}

	return zeroes
}

func getSteps(filename string) []int {
	steps := []int{}
	for _, line := range shared.ReadFile(filename) {
		step, _ := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(line, "L", "-"), "R", ""))
		steps = append(steps, step)
	}
	return steps
}

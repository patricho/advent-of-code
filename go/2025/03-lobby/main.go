package main

import (
	"strconv"

	"github.com/fatih/color"
	"github.com/patricho/advent-of-code/go/shared"
)

const (
	INPUT_FILE = "../inputs/2025/03-input.txt"
	TEST_FILE  = "../inputs/2025/03-test.txt"
)

func main() {
	shared.DisplayMain(func() {
		test()
		run()
	})
}

func test() {
	shared.RunCase("test part 1", func() int { return part1(TEST_FILE) }, 357)
	shared.RunCase("test part 2", func() int { return part2(TEST_FILE) }, 3121910778619)
}

func run() {
	shared.RunCase("part 1", func() int { return part1(INPUT_FILE) }, 17229)
	shared.RunCase("part 2", func() int { return part2(INPUT_FILE) }, 170520923035051)
}

func part1(filename string) int {
	return checkFile(filename, 2)
}

func part2(filename string) int {
	return checkFile(filename, 12)
}

func checkFile(filename string, targetLen int) int {
	input := shared.ReadFile(filename)

	result := 0
	checks := 0

	for _, line := range input {
		n, c := checkRange(line, targetLen)
		result += n
		checks += c
	}

	color.Blue("checks: %d", checks)

	return result
}

func checkRange(row string, targetLen int) (int, int) {
	startIdx := 0
	checks := 0
	result := ""

	maxIdx := len(row) - targetLen

	for {
		pick := '0'

		for i := startIdx; i <= maxIdx; i++ {
			candidate := rune(row[i])
			checks++

			if candidate > pick {
				pick = candidate
				startIdx = i + 1
			}
		}

		result += string(pick)

		if len(result) >= targetLen {
			break
		}

		maxIdx++
	}

	n, _ := strconv.Atoi(result)
	return n, checks
}

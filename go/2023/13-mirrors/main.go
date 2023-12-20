package main

import (
	"fmt"
	"slices"

	"github.com/patricho/advent-of-code/go/util"
)

var (
	blocks           [][]string
	lines            []string
	binaryLines      []uint64
	wantDiffs        int
	pointsMultiplier int
)

func main() {
	run(1, "test1.txt")
	run(2, "test1.txt")

	run(1, "input.txt")
	run(2, "input.txt")
}

func run(part int, filename string) {
	alllines := util.ReadFile(filename)
	blocks = findBlocks(alllines)
	sum := 0

	for i, block := range blocks {
		lines = block
		sum += processBlock(part, i)
		// fmt.Println("")
	}

	fmt.Println(part, filename, "total sum", sum)
}

func processBlock(part, num int) int {
	binaryLines = util.ToBinary(lines, "#", ".")
	wantDiffs = 0
	pointsMultiplier = 100

	if part == 2 {
		wantDiffs = 1
	}

	points := calculatePoints()
	if points > 0 {
		// found horizontal match
		// fmt.Println(num, "horizontal points", points)
		return points
	}

	// fmt.Println("no horizontal match...")

	// check vertical too
	lines = flipLines(lines)
	binaryLines = util.ToBinary(lines, "#", ".")
	pointsMultiplier = 1

	points = calculatePoints()
	if points > 0 {
		// found vertical match
		// fmt.Println(num, "vertical points", points)
		return points
	}

	// // if we get this far, something is fucked
	// // then print out the current block for debug purposes
	for _, r := range lines {
		fmt.Println(num, "block err", r)
	}

	panic("wtf")
}

func findBlocks(lines []string) [][]string {
	blocks := [][]string{}
	block := []string{}

	for i := range lines {
		if i < len(lines)-1 && len(lines[i]) == 0 {
			blocks = append(blocks, block)
			block = []string{}
			continue
		}

		if len(lines[i]) > 0 {
			block = append(block, lines[i])
			continue
		}
	}

	if len(block) > 0 {
		blocks = append(blocks, block)
	}

	return blocks
}

func calculatePoints() int {
	// check each line with the one before, from line 2-n
	for i := 1; i < len(binaryLines); i++ {
		// fmt.Println("debug binarylines", i, binaryLines)

		line := binaryLines[i]
		prevline := binaryLines[i-1]

		// if match found
		if util.CountDiffs(line, prevline) <= wantDiffs {
			// fmt.Println("match found on line", i-1, "and", i)

			// fmt.Println("all binary lines", binaryLines)

			// NOTE: make sure to CLONE from binarylines, since the
			// slices.reverse below changes IN PLACE, this also affects the
			// original binarylines slice...
			linesBefore := slices.Clone(binaryLines[:i])
			linesAfter := binaryLines[i:]

			// fmt.Println("binary lines before", linesBefore)
			// fmt.Println("binary lines after", linesAfter)

			slices.Reverse(linesBefore)

			// fmt.Println("reversed binary lines before", linesBefore)

			maxLen := len(linesBefore)
			if len(linesAfter) < maxLen {
				maxLen = len(linesAfter)
			}

			// for _, l := range lines[:i] {
			// 	// fmt.Println("lines before", linesBefore, l)
			// }
			// for _, l := range lines[i:] {
			// 	// fmt.Println("lines after ", linesAfter, l)
			// }

			// check if the match is valid for the entire array
			if isMatchValid(linesBefore, linesAfter, maxLen) {
				// fmt.Println("valid match found on line", i-1, "and", i)
				points := len(linesBefore) * pointsMultiplier
				return points
			}
		}
	}

	return -1
}

func isMatchValid(linesBefore, linesAfter []uint64, maxLen int) bool {
	totalDiffs := 0

	for i := 0; i < maxLen; i++ {
		totalDiffs += util.CountDiffs(linesBefore[i], linesAfter[i])
	}

	return totalDiffs == wantDiffs
}

func flipLines(lines []string) []string {
	flipped := make([]string, len(lines[0]))
	for col := 0; col < len(lines[0]); col++ {
		flipped[col] = ""
		for row := 0; row < len(lines); row++ {
			flipped[col] += string(lines[row][col])
		}
	}
	return flipped
}

package treetoptreehouse08

import (
	"fmt"
	"sort"
	"strconv"

	c "github.com/patricho/advent-of-code/go/util"
)

func Part1() {
	rows := c.ReadFile("08-treetop-tree-house/input.txt")

	trees := splitData(rows)

	fmt.Println("")

	visible := [][]string{}
	visibleCount := 0

	for row, rowval := range trees {
		cols := []string{}
		for col, colval := range rowval {
			vis := checkVisibility(trees, row, col)

			if vis {
				cols = append(cols, colval)
				visibleCount++
			} else {
				cols = append(cols, "_")
			}
		}

		visible = append(visible, cols)
	}

	for _, row := range visible {
		fmt.Println(row)
	}

	fmt.Println("visibleCount", visibleCount)
}

func Part2() {
	rows := c.ReadFile("08-treetop-tree-house/input.txt")

	trees := splitData(rows)

	fmt.Println("")

	visibleScores := []int{}

	for row, rowval := range trees {
		for col := range rowval {
			visibleScores = append(visibleScores, countVisibility(trees, row, col))
		}
	}

	fmt.Println("")

	sort.Ints(visibleScores)

	for _, score := range visibleScores {
		fmt.Println("score", score)
	}
}

func checkVisibility(trees [][]string, row, col int) bool {
	height, _ := strconv.Atoi(trees[row][col])

	allAboveVisible, _ := checkAbove(trees, height, row, col)
	allBelowVisible, _ := checkBelow(trees, height, row, col)
	allLeftVisible, _ := checkLeft(trees, height, row, col)
	allRightVisible, _ := checkRight(trees, height, row, col)

	return allAboveVisible || allBelowVisible || allLeftVisible || allRightVisible
}

func countVisibility(trees [][]string, row, col int) int {
	height, _ := strconv.Atoi(trees[row][col])

	_, scoreAbove := checkAbove(trees, height, row, col)
	_, scoreBelow := checkBelow(trees, height, row, col)
	_, scoreLeft := checkLeft(trees, height, row, col)
	_, scoreRight := checkRight(trees, height, row, col)

	fmt.Println("scores", scoreAbove, scoreLeft, scoreRight, scoreBelow)

	return scoreAbove * scoreBelow * scoreLeft * scoreRight
}

func checkLeft(trees [][]string, height, row, col int) (bool, int) {
	allLeftVisible := true
	leftVisibleCount := 0

	for c := col - 1; c >= 0; c-- {
		neighborHeight, _ := strconv.Atoi(trees[row][c])

		// fmt.Println("left", row, c, "height", neighborHeight, "visible!")
		leftVisibleCount++

		if neighborHeight >= height {
			allLeftVisible = false
			break
		}
	}

	// fmt.Println("left return", leftVisibleCount)

	return allLeftVisible, leftVisibleCount
}

func checkRight(trees [][]string, height, row, col int) (bool, int) {
	allRightVisible := true
	rightVisibleCount := 0

	for c := col + 1; c < len(trees[row]); c++ {
		neighborHeight, _ := strconv.Atoi(trees[row][c])

		// fmt.Println("right", row, c, "height", neighborHeight, "visible!")
		rightVisibleCount++

		if neighborHeight >= height {
			allRightVisible = false
			break
		}
	}

	// fmt.Println("right return", rightVisibleCount)

	return allRightVisible, rightVisibleCount
}

func checkAbove(trees [][]string, height, row, col int) (bool, int) {
	allAboveVisible := true
	aboveVisibleCount := 0

	for i := row - 1; i >= 0; i-- {
		neighborHeight, _ := strconv.Atoi(trees[i][col])

		// fmt.Println("above", i, col, "height", neighborHeight, "visible!")
		aboveVisibleCount++

		if neighborHeight >= height {
			allAboveVisible = false
			break
		}
	}

	// fmt.Println("above return", aboveVisibleCount)

	return allAboveVisible, aboveVisibleCount
}

func checkBelow(trees [][]string, height, row, col int) (bool, int) {
	allBelowVisible := true
	belowVisibleCount := 0

	for i := row + 1; i < len(trees); i++ {
		neighborHeight, _ := strconv.Atoi(trees[i][col])

		// fmt.Println("below", i, col, "height", neighborHeight, "visible!")
		belowVisibleCount++

		if neighborHeight >= height {
			allBelowVisible = false
			break
		}
	}

	// fmt.Println("below return", belowVisibleCount)

	return allBelowVisible, belowVisibleCount
}

func splitData(rows []string) [][]string {
	trees := [][]string{}

	for _, row := range rows {
		cols := []string{}

		for _, col := range row {
			cols = append(cols, string(col))
		}

		trees = append(trees, cols)

		fmt.Println(cols)
	}

	return trees
}

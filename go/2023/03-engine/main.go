package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/patricho/advent-of-code/go/util"
)

var (
	lines   []string
	engines map[string][]int
)

func getLinePart(idx, start, end int) string {
	if idx < 0 || idx >= len(lines) {
		return ""
	}
	return lines[idx][start:end]
}

func checkEngine(line string, idx, start, num int) {
	asteriskRegex, _ := regexp.Compile(`\*`)
	for _, asxIdxs := range asteriskRegex.FindAllStringIndex(line, -1) {
		asxLineIdx := start + asxIdxs[0]
		key := fmt.Sprint(idx, "x", asxLineIdx)
		engines[key] = append(engines[key], num)
	}
}

func main() {
	lines = util.ReadFile("input.txt")
	numbersRegex, _ := regexp.Compile(`\d+`)
	symbolsRegex, _ := regexp.Compile(`[^\d.]`)
	sumPart1 := 0
	sumPart2 := 0
	engines = map[string][]int{}

	for lineIdx, line := range lines {
		for _, numIdxs := range numbersRegex.FindAllStringIndex(line, -1) {
			numLen := numIdxs[1] - numIdxs[0]
			num, _ := strconv.Atoi(line[numIdxs[0] : numIdxs[0]+numLen])

			start := numIdxs[0] - 1
			if start < 0 {
				start = 0
			}

			end := numIdxs[0] + numLen + 1
			if end > len(line) {
				end = len(line)
			}

			above := getLinePart(lineIdx-1, start, end)
			this := getLinePart(lineIdx, start, end)
			below := getLinePart(lineIdx+1, start, end)

			if symbolsRegex.MatchString(fmt.Sprint(above, this, below)) {
				sumPart1 += num
			}

			checkEngine(above, lineIdx-1, start, num)
			checkEngine(this, lineIdx, start, num)
			checkEngine(below, lineIdx+1, start, num)
		}
	}

	fmt.Println(sumPart1)

	for _, nums := range engines {
		if len(nums) == 2 {
			sumPart2 += nums[0] * nums[1]
		}
	}

	fmt.Println(sumPart2)
}

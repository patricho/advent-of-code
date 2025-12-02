package main

import (
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/fatih/color"

	"github.com/patricho/advent-of-code/go/shared"
)

var (
	reTwice    *regexp2.Regexp
	reMultiple *regexp2.Regexp
)

func main() {
	// Start, a sequence of digits, then backtrack to match that sequence again, end
	reTwice = regexp2.MustCompile(`^(\d+)\1$`, 0)

	// Start, a sequence of digits, then backtrack to match that sequence again one or more times, end
	reMultiple = regexp2.MustCompile(`^(\d+)\1+$`, 0)

	test()
	run()

	color.RGB(64, 64, 64).Printf("\n---\n")
}

func test() {
	shared.RunCase("test part 1", func() int {
		return part1("../inputs/2025/02-test.txt")
	}, 1227775554)

	shared.RunCase("test part 2", func() int {
		return part2("../inputs/2025/02-test.txt")
	}, 4174379265)
}

func run() {
	shared.RunCase("part 1", func() int {
		return part1("../inputs/2025/02-input.txt")
	}, 12586854255)

	shared.RunCase("part 2", func() int {
		return part2("../inputs/2025/02-input.txt")
	}, 17298174201)
}

func part1(filename string) int {
	return check(filename, reTwice)
}

func part2(filename string) int {
	return check(filename, reMultiple)
}

func check(filename string, re *regexp2.Regexp) int {
	input := shared.ReadFile(filename)
	errsum := 0
	ranges := strings.Split(input[0], ",")
	for _, rng := range ranges {
		errsum += checkRange(rng, re)
	}
	return errsum
}

func checkRange(rng string, re *regexp2.Regexp) int {
	errsum := 0

	arr := strings.Split(rng, "-")

	start, _ := strconv.Atoi(arr[0])
	end, _ := strconv.Atoi(arr[1])

	for n := start; n <= end; n++ {
		nstr := strconv.Itoa(n)
		if ok, _ := re.MatchString(nstr); ok {
			errsum += n
		}
	}

	return errsum
}

package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/patricho/advent-of-code/go/util"
)

func main() {
	// filename := os.Args[1]
	part2()
}

func part1() {
	lines := util.ReadFile("input.txt")
	digits := "0123456789"
	sum := 0

	for _, line := range lines {
		digitsInLine := []string{}

		for _, char := range line {
			if !strings.Contains(digits, string(char)) {
				continue
			}
			digitsInLine = append(digitsInLine, string(char))
		}

		first := digitsInLine[0]
		last := digitsInLine[len(digitsInLine)-1]
		twochars := fmt.Sprint(first, last)
		twodigits, _ := strconv.Atoi(twochars)
		sum += twodigits
	}

	fmt.Println(sum)
}

func part2() {
	lines := util.ReadFile("input.txt")
	digNumbers := "0123456789"
	letNumbers := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	sum := 0

	for _, line := range lines {
		numbersInLine := []string{}

		for idx, char := range line {
			if strings.Contains(digNumbers, string(char)) {
				numbersInLine = append(numbersInLine, string(char))
				continue
			}

			restOfLine := line[idx:]

			for wordNum, num := range letNumbers {
				if strings.HasPrefix(restOfLine, wordNum) {
					numbersInLine = append(numbersInLine, num)
				}
			}
		}

		first := numbersInLine[0]
		last := numbersInLine[len(numbersInLine)-1]
		twoCharNumString := fmt.Sprint(first, last)
		twoCharNumber, _ := strconv.Atoi(twoCharNumString)
		sum += twoCharNumber
	}

	fmt.Println(sum)
}

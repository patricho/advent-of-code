package main

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strings"

	"github.com/patricho/advent-of-code/go/shared"
)

func main() {
	lines := shared.ReadFile("input.txt")
	cardCounts := map[int]int{}
	numbersRegex, _ := regexp.Compile(`\d+`)
	scoreSum := 0
	cardsSum := 0

	for idx, line := range lines {
		numberSeries := strings.Split(line[strings.Index(line, ":")+1:], "|")
		winningNumbers := numbersRegex.FindAllString(numberSeries[0], -1)
		hasNumbers := numbersRegex.FindAllString(numberSeries[1], -1)

		cardCounts[idx] += 1
		matches := 0

		for _, n := range hasNumbers {
			if slices.Contains(winningNumbers, n) {
				matches++
			}
		}

		scoreSum += int(math.Pow(2, float64(matches-1)))

		for i := 1; i <= matches; i++ {
			cardCounts[idx+i] += cardCounts[idx]
		}

		cardsSum += cardCounts[idx]
	}

	fmt.Println("score:", scoreSum)
	fmt.Println("cards:", cardsSum)
}

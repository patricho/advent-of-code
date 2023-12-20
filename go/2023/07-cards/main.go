package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/patricho/advent-of-code/go/shared"
)

type Hand struct {
	Hand         string
	SortableHand string
	TypeScores   []int
	TypeScore    int
	Bid          int
	Rank         int
	Winnings     int
}

func main() {
	hands := parseHands()
	run(hands, 1)
	run(hands, 2)
}

func run(hands []Hand, part int) {
	shared.Measure(func() {
		rankHands(&hands, part)
		sum := 0
		for _, h := range hands {
			sum += h.Winnings
		}
		fmt.Println("winnings:", sum)
	})
}

func parseHands() []Hand {
	out := []Hand{}
	lines := shared.ReadFile("input.txt")
	for _, line := range lines {
		arr := strings.Split(line, " ")
		out = append(out, Hand{
			Hand: arr[0],
			Bid:  shared.ToInt(arr[1]),
		})
	}
	return out
}

func getTypeScores(h Hand, part int) ([]int, int) {
	if part == 1 {
		return getTypeScoresPart1(h)
	}
	return getTypeScoresPart2(h)
}

func getTypeScoresPart2(h Hand) ([]int, int) {
	typeScores := []int{}
	score := 0

	// do one count for each card type and pick the best one
	for _, rff := range "23456789TQKA" {
		jokerHand := strings.ReplaceAll(h.Hand, "J", string(rff))
		jokerTypeScores, jokerScore := calculateTypeScores(jokerHand)
		if jokerScore > score {
			typeScores = jokerTypeScores
			score = jokerScore
		}
	}

	return typeScores, score
}

func getTypeScoresPart1(h Hand) ([]int, int) {
	return calculateTypeScores(h.Hand)
}

func calculateTypeScores(hand string) ([]int, int) {
	typeScores := []int{}
	typeScore := 0

	for _, rf := range hand {
		count := 0

		for _, rc := range hand {
			if rc == rf {
				count++
			}
		}

		typeScores = append(typeScores, count)
		typeScore += count
	}

	slices.Sort(typeScores)
	slices.Reverse(typeScores)

	return typeScores, typeScore
}

func rankHands(hands *[]Hand, part int) {
	// calculate type score and sort cards by strength
	for idx, h := range *hands {
		typeDef, typeScore := getTypeScores(h, part)
		sortableHand := getSortableHand(h, part)

		(*hands)[idx].TypeScores = typeDef
		(*hands)[idx].TypeScore = typeScore
		(*hands)[idx].SortableHand = sortableHand
	}

	// sort list first by type score and then strength
	slices.SortFunc(*hands, func(a Hand, b Hand) int {
		at := a.TypeScore
		bt := b.TypeScore

		ah := a.SortableHand
		bh := b.SortableHand

		if at < bt {
			return 1
		} else if at == bt {
			if ah < bh {
				return 1
			} else if ah == bh {
				return 0
			} else {
				return -1
			}
		} else {
			return -1
		}
	})

	// rank and calculate winnings based on placement in list
	for i, h := range *hands {
		rank := len(*hands) - i
		(*hands)[i].Rank = rank
		(*hands)[i].Winnings = rank * h.Bid
	}
}

// getSortableHand translates card keys to an lexicographically sortable string
func getSortableHand(h Hand, part int) string {
	charsFrom := "23456789TJQKA"
	strengthsTo := "0123456789ABC"

	if part == 2 {
		// jokers are now the weakest card
		charsFrom = "J23456789TQKA"
	}

	translatedHand := ""
	for _, r := range h.Hand {
		idxFrom := strings.Index(charsFrom, string(r))
		translatedHand += string(strengthsTo[idxFrom])
	}
	return translatedHand
}

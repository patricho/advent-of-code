package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/patricho/advent-of-code/go/util"
)

func main() {
	lines := util.ReadFile("input.txt")
	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	part1Sum := 0
	part2Sum := 0

	for _, line := range lines {
		lineparts := strings.Split(line, ": ")
		gameid, _ := strconv.Atoi(strings.Split(lineparts[0], " ")[1])
		games := strings.Split(lineparts[1], "; ")
		lineok := true
		minPerColor := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, game := range games {
			gameparts := strings.Split(game, ", ")

			for _, gamepart := range gameparts {
				colorandcount := strings.Split(gamepart, " ")
				color := colorandcount[1]
				count, _ := strconv.Atoi(colorandcount[0])
				if count > limits[color] {
					lineok = false
				}

				if count > minPerColor[color] {
					minPerColor[color] = count
				}
			}
		}

		pow := 1
		for _, val := range minPerColor {
			pow *= val
		}
		part2Sum += pow

		fmt.Println("game", gameid, "minpercolor", minPerColor, "pow", pow)

		if lineok {
			part1Sum += gameid
		}
	}

	fmt.Println("part1:", part1Sum)
	fmt.Println("part2:", part2Sum)
}

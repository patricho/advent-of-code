package cathoderaytube10

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	c "github.com/patricho/advent-of-code/go/shared"
)

type Instruction struct {
	noop   bool
	value  int
	cycles int
}

func Part1() {
	instructionStrings := c.ReadFile("10-cathode-ray-tube/input.txt")

	instructions := parse(instructionStrings)

	X := 1
	cycle := 1
	sum := 0

	for _, i := range instructions {
		if i.noop {
			// fmt.Println(cycle, "noop start", X)
			draw(cycle, X)
			cycle++
			checkSignalStrength(&sum, cycle, X)
			// fmt.Println(cycle, "noop done", X)
		} else {
			// fmt.Println(cycle, "addx start", X)
			draw(cycle, X)
			cycle++
			checkSignalStrength(&sum, cycle, X)
			draw(cycle, X)
			cycle++
			X += i.value
			// fmt.Println(cycle, "addx done", X)
			checkSignalStrength(&sum, cycle, X)
		}
	}

	fmt.Println("sum", sum)
}

func draw(cycle, X int) {
	start, end := X-1, X+1

	idx := cycle - 1

	pixel := (idx - int(math.Floor(float64(idx/40))*40))

	if start <= pixel && end >= pixel {
		// fmt.Println("cycle", cycle, "pixel", pixel, "X", X, "draw #")
		fmt.Print("#")
	} else {
		// fmt.Println("cycle", cycle, "pixel", pixel, "X", X, "draw .", "cycle / 40", math.Floor(float64(cycle/40)))
		fmt.Print(".")
	}

	if (cycle)%40 == 0 {
		fmt.Print("\n")
	}
}

func parse(instructionStrings []string) []Instruction {
	instructions := []Instruction{}

	for _, str := range instructionStrings {
		if str == "noop" {
			instructions = append(instructions, Instruction{noop: true, cycles: 1})
		} else {
			value, _ := strconv.Atoi(strings.Split(str, " ")[1])
			instructions = append(instructions, Instruction{noop: false, value: value, cycles: 2})
		}
	}

	return instructions
}

func checkSignalStrength(sum *int, cycle int, X int) {
	if (cycle-20)%40 == 0 {
		// fmt.Println("check cycle", cycle, "X", X, "score", cycle*X)
		*sum += cycle * X
	}
}

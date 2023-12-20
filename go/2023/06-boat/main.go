package main

import "fmt"

func main() {
	testraces := [][]int64{
		{7, 9},
		{15, 40},
		{30, 200},
	}

	races := [][]int64{
		{49, 263},
		{97, 1_532},
		{94, 1_378},
		{94, 1_851},
	}

	part2test := [][]int64{
		{71_530, 940_200},
	}

	part2races := [][]int64{
		{49_979_494, 263_153_213_781_851},
	}

	solve(testraces)
	solve(races)
	solve(part2test)
	solve(part2races)
}

func solve(races [][]int64) {
	sum := 1

	for _, race := range races {
		time := race[0]
		record := race[1]
		wins := 0

		for t := int64(0); t <= time; t++ {
			timeleft := time - t
			speed := t
			distance := timeleft * speed
			win := distance > record
			if win {
				wins++
			}

			// fmt.Println("elapsed", t, "distance", distance, "win", win)

			if !win && wins > 0 {
				break
			}
		}

		sum *= wins

		fmt.Println("race", race, "wins", wins)
	}

	fmt.Println("sum", sum)
}

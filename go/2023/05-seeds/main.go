package main

import (
	"fmt"
	"strings"

	"github.com/patricho/advent-of-code/go/util"
)

type Pair struct {
	From int64
	To   int64
}

type MappingLegend struct {
	From  int64
	To    int64
	Range int64
}

var (
	legendSets   [][]MappingLegend
	seeds        []int64
	lowestResult int64
)

func parseLegendSets() {
	lines := util.ReadFile("input.txt")

	legendSets = [][]MappingLegend{}
	legendSet := []MappingLegend{}

	for idx, line := range lines {
		if idx == 0 {
			// read seeds
			nums := strings.Split(strings.TrimPrefix(line, "seeds: "), " ")
			for _, num := range nums {
				seeds = append(seeds, util.ToInt64(num))
			}
		} else if idx >= 2 {
			// read legends
			if len(line) == 0 {
				// new set
				legendSets = append(legendSets, legendSet)
				legendSet = []MappingLegend{}
			} else if strings.Contains(line, "map") {
				// skip name
			} else {
				// parse an ordinary damn line
				nums := strings.Split(line, " ")
				legendSet = append(legendSet, MappingLegend{
					To:    util.ToInt64(nums[0]),
					From:  util.ToInt64(nums[1]),
					Range: util.ToInt64(nums[2]),
				})
			}
		}
	}

	if len(legendSet) > 0 {
		legendSets = append(legendSets, legendSet)
	}
}

func main() {
	parseLegendSets()
	// part1()
	part2()
	// ext(util.ReadFile("input.txt"))
}

func checkSeed(seed int64) {
	result := seed

	for _, legendSet := range legendSets {
		for _, legend := range legendSet {
			start := legend.From
			end := legend.From + legend.Range
			offset := legend.To - legend.From

			if result >= start && result < end {
				result += offset
				break
			}
		}
	}

	if lowestResult < 0 || result < lowestResult {
		lowestResult = result
	}
}

func part1() {
	lowestResult = -1

	for _, seed := range seeds {
		checkSeed(seed)
	}

	fmt.Println("part 1 lowest result", lowestResult)
}

func part2() {
	lowestResult = -1

	for i := 0; i < len(seeds); i += 2 {
		seedStart := seeds[i]
		seedRange := seeds[i+1]
		seedEnd := seedStart + seedRange - 1

		fmt.Println("part2", i, "start", seedStart, "range", seedRange, "end", seedEnd)

		for seed := seedStart; seed <= seedEnd; seed++ {
			checkSeed(seed)
		}

		fmt.Println("part2", i, "lowest result", lowestResult)
	}

	fmt.Println("part 2 lowest result", lowestResult)
}

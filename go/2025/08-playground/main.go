package main

import (
	"maps"
	"math"
	"slices"
	"strconv"

	s "github.com/patricho/advent-of-code/go/shared"
)

const (
	INPUT_FILE = "../inputs/2025/08-input.txt"
	TEST_FILE  = "../inputs/2025/08-test.txt"
)

type Pair struct {
	a string
	b string
	d float64
}

type Junction struct {
	p s.Point
	c string
}

var (
	coords   map[string]Junction
	pairs    []Pair
	circuits map[string]int
)

func main() {
	s.DisplayMain(func() {
		test()
		run()
	})
}

func test() {
	s.RunCase("test part 1", func() int { return part1(TEST_FILE, 10) }, 40)
	s.RunCase("test part 2", func() int { return part2(TEST_FILE) }, 25272)
}

func run() {
	s.RunCase("part 1", func() int { return part1(INPUT_FILE, 1000) }, 244188)
	s.RunCase("part 2", func() int { return part2(INPUT_FILE) }, 8361881885)
}

func part1(filename string, target int) int {
	return solve(filename, 1, target)
}

func part2(filename string) int {
	return solve(filename, 2, 0)
}

func solve(filename string, part int, target int) int {
	parseInput(filename)

	circuitNo := 0
	connected := map[string]bool{}
	coordsNo := len(coords)
	latest := [2]int{0, 0}

	for _, pair := range pairs {
		if connected[pair.a+pair.b] {
			continue
		}

		circuitNo++

		if len(coords[pair.a].c) > 0 && len(coords[pair.b].c) > 0 &&
			coords[pair.a].c != coords[pair.b].c {
			// Merge related circuits
			tomerge := []string{coords[pair.a].c, coords[pair.b].c}
			for k, c := range coords {
				if slices.Contains(tomerge, c.c) {
					c.c = coords[pair.a].c
					coords[k] = c
				}
			}
		} else {
			nextcircuit := coords[pair.a].c
			if len(nextcircuit) == 0 {
				nextcircuit = coords[pair.b].c
			}
			if len(nextcircuit) == 0 {
				nextcircuit = strconv.Itoa(circuitNo)
			}

			a := coords[pair.a]
			b := coords[pair.b]

			a.c = nextcircuit
			b.c = nextcircuit

			coords[pair.a] = a
			coords[pair.b] = b
		}

		connected[pair.a+pair.b] = true
		connected[pair.b+pair.a] = true

		if part == 1 && circuitNo >= target {
			break
		}

		if part == 2 {
			countCircuits()
			if len(circuits) == 1 {
				// Check if all coordinates are in this circuit
				firstCircuitKey := ""
				for k := range circuits {
					firstCircuitKey = k
					break
				}
				if circuits[firstCircuitKey] == coordsNo {
					latest = [2]int{coords[pair.a].p.X, coords[pair.b].p.X}
					break
				}
			}
		}
	}

	if part == 1 {
		countCircuits()

		values := slices.Collect(maps.Values(circuits))
		slices.SortFunc(values, func(a int, b int) int {
			return b - a
		})

		sum := 1
		for i := range 3 {
			sum *= values[i]
		}

		return sum
	}

	return latest[0] * latest[1]
}

func countCircuits() {
	circuits = map[string]int{}

	for _, c := range coords {
		if len(c.c) == 0 {
			continue
		}
		circuits[c.c] = circuits[c.c] + 1
	}
}

func parseInput(filename string) {
	coords = map[string]Junction{}
	pairs = []Pair{}

	lines := s.ReadFile(filename)

	for _, line := range lines {
		nums := s.ToIntSlice(line, ",")
		p := s.Point{X: nums[0], Y: nums[1], Z: nums[2]}
		coords[line] = Junction{p: p}
	}

	for ka, ca := range coords {
		for kb, cb := range coords {
			if ka == kb {
				continue
			}

			dx := ca.p.X - cb.p.X
			dy := ca.p.Y - cb.p.Y
			dz := ca.p.Z - cb.p.Z

			dist := math.Sqrt(float64(dx*dx + dy*dy + dz*dz))

			pairs = append(pairs, Pair{a: ka, b: kb, d: dist})
		}
	}

	slices.SortFunc(pairs, func(a Pair, b Pair) int {
		return int(a.d - b.d)
	})
}

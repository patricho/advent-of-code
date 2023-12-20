package shared

import (
	"fmt"
	"testing"
)

// primitiveGridToString is a crude KISS string concatenator,
// to have something to compare the optimized function to
func primitiveGridToString(grid [][]rune) string {
	out := ""
	for _, line := range grid {
		for _, r := range line {
			out += string(r)
		}
	}
	return out
}

func BenchmarkPrimitiveGridToString(b *testing.B) {
	lines := ReadFile("../14-rocks/input.txt")
	grid := LinesToGrid(lines)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		primitiveGridToString(grid)
	}
}

func BenchmarkGridToString(b *testing.B) {
	lines := ReadFile("../14-rocks/input.txt")
	grid := LinesToGrid(lines)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		GridToString(grid)
	}
}

func BenchmarkHashString(b *testing.B) {
	lines := ReadFile("../14-rocks/input.txt")
	grid := LinesToGrid(lines)
	gridstring := GridToString(grid)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		HashString(gridstring)
	}
}

var countDiffsTestData = []struct {
	a  uint64
	b  uint64
	as string
	bs string
}{
	{a: 281, b: 281, as: "#...##..#", bs: "#...##..#"},
	{a: 57063, b: 82543, as: ".##.####.###..###", bs: "#.#....#..##.####"},
	{a: 1923580808021856, b: 1829883649014368, as: ".##.##.#.#.#.#####..#...###.##....#..#...###.##.....", bs: ".##.#........#...#..######..###....#..###.#..##....."},
}

func BenchmarkCountDiffsInt(b *testing.B) {
	for _, v := range countDiffsTestData {
		b.Run(fmt.Sprintf("countDiffsInt(%d, %d)", v.a, v.b), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CountDiffs(v.a, v.b)
			}
		})
	}
}

func BenchmarkCountDiffsString(b *testing.B) {
	for _, v := range countDiffsTestData {
		b.Run(fmt.Sprintf("countDiffsString(%s, %s)", v.as, v.bs), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				countDiffsString(v.as, v.bs)
			}
		})
	}
}

// countDiffsString is an inefficient way to check, just to compare with countDiffsInt
func countDiffsString(a, b string) int {
	n := 0
	for i := range a {
		if a[i] != b[i] {
			n++
		}
	}
	return n
}

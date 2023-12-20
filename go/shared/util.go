package shared

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"math"
	"math/bits"
	"os"
	"strconv"
	"strings"
	"time"
)

// Abs returns the absolute value of x
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ToInt(input string) int {
	res, _ := strconv.Atoi(input)
	return res
}

func ToInt64(input string) int64 {
	res, _ := strconv.Atoi(input)
	return int64(res)
}

func ReadFile(file string) []string {
	readFile, _ := os.Open(file)
	defer readFile.Close()

	lines := []string{}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines
}

func LinesToGrid(lines []string) [][]rune {
	out := make([][]rune, len(lines))
	for lidx, line := range lines {
		out[lidx] = make([]rune, len(line))
		for i, r := range line {
			out[lidx][i] = r
		}
	}
	return out
}

func HashString(in string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(in)))
}

func GridToString(grid [][]rune) string {
	var b strings.Builder
	b.Grow(len(grid) * len(grid[0]))
	for _, line := range grid {
		for _, r := range line {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func OOB[T any](grid [][]T, p Point) bool {
	return p.Y < 0 || p.Y >= len(grid) || p.X < 0 || p.X >= len(grid[0])
}

func Measure(f func()) {
	start := time.Now()
	f()
	fmt.Println("time:", time.Since(start))
}

// GCD finds greatest common divisor via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM finds least common multiple via GCD
func LCM(integers ...int) int {
	a := integers[0]
	b := integers[1]

	result := a * b / GCD(a, b)

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func ManhattanDistance(a []int, b []int) float64 {
	var n int
	var s float64

	if len(a) != len(b) {
		fmt.Println("Be sure that both vectors have the same dimension!")
		return 0
	}

	n = len(a)

	s = 0

	for i := 0; i < n; i += 1 {
		s += math.Abs(float64(b[i]) - float64(a[i]))
	}

	return float64(s)
}

// ToBinaryString takes an int and just returns it as a print friendly string
// ie 1001 -> "1001"
func ToBinaryString(in uint64) string {
	return strconv.FormatInt(int64(in), 2)
}

// CountDiffs does a binary XOR between A and B and counts the number of differences
func CountDiffs(a, b uint64) int {
	return bits.OnesCount64(a ^ b)
}

// ToBinary takes in a string slice and converts it to an int slice
// so for example ...#..#.# becomes 000100101 becomes 37
func ToBinary(lines []string, one string, zero string) []uint64 {
	out := []uint64{}
	for _, l := range lines {
		bl := strings.ReplaceAll(strings.ReplaceAll(l, one, "1"), zero, "0")
		b, _ := strconv.ParseInt(bl, 2, 64)
		out = append(out, uint64(b))
	}
	return out
}

func Last[T any](slice []T) T {
	return slice[len(slice)-1]
}

func All[T comparable](slice []T, find T) bool {
	for _, n := range slice {
		if n != find {
			return false
		}
	}

	return true
}

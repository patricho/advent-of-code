package main

import (
	"fmt"
	"strings"

	"github.com/patricho/advent-of-code/go/util"
)

type Lens struct {
	Label string
	Hash  int
	Focal int
}

var boxes map[int][]Lens

func main() {
	lines := util.ReadFile("input.txt")
	lenses := strings.Split(lines[0], ",")
	runPart1(lenses)
	runPart2(lenses)
}

func runPart2(lenses []string) {
	boxes = make(map[int][]Lens, 256)
	for i := 0; i < 256; i++ {
		boxes[i] = []Lens{}
	}

	for _, lensstr := range lenses {
		add := true
		if strings.Contains(lensstr, "-") {
			add = false
			lensstr = strings.ReplaceAll(lensstr, "-", "=")
		}

		arr := strings.Split(lensstr, "=")
		lens := Lens{
			Label: arr[0],
			Hash:  lensHash(arr[0]),
			Focal: util.ToInt(arr[1]),
		}

		if add {
			addLens(lens)
		} else {
			removeLens(lens)
		}
	}

	sum := 0

	for key, box := range boxes {
		for li, lens := range box {
			power := (key + 1) * (li + 1) * lens.Focal
			sum += power
		}
	}

	fmt.Println("part 2 sum:", sum)
}

func removeLens(lens Lens) {
	box := boxes[lens.Hash]
	newbox := []Lens{}

	for _, l := range box {
		if l.Label != lens.Label {
			newbox = append(newbox, l)
		}
	}

	boxes[lens.Hash] = newbox
}

func addLens(lens Lens) {
	box := boxes[lens.Hash]
	found := false
	for i, l := range box {
		if l.Label == lens.Label {
			box[i] = lens
			found = true
		}
	}

	if !found {
		box = append(box, lens)
	}

	boxes[lens.Hash] = box
}

func runPart1(lenses []string) {
	sum := 0
	for _, lens := range lenses {
		sum += lensHash(lens)
		// fmt.Println(lens, sum)
	}
	fmt.Println("part 1 sum:", sum)
}

func lensHash(lens string) int {
	sum := 0
	for _, r := range lens {
		sum += int(r)
		sum *= 17
		sum = sum % 256
	}
	return sum
}

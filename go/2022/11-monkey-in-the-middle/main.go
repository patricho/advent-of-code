package main

import (
	"fmt"
	"math/big"
	"sort"
)

type Monkey struct {
	id          int
	items       []*big.Int
	inspections int
	op          func(*big.Int)
	test        func(*big.Int) bool
	truemon     *Monkey
	falsemon    *Monkey
}

func main() {
	Part1()
}

func Part1() {
	p, monkeys := createMonkeys()

	for round := 1; round <= 20; round++ {
		for _, monkey := range monkeys {
			throwItems(monkey, 1, p)
		}
	}

	sums := []int{}

	for _, monkey := range monkeys {
		fmt.Println("monkey", monkey.id, "items", (*monkey).items, "inspections", monkey.inspections)
		sums = append(sums, monkey.inspections)
	}

	sort.Ints(sums)
	sum := sums[len(sums)-1] * sums[len(sums)-2]

	fmt.Println("sum", sum)
}

func Part2() {
	p, monkeys := createMonkeys()

	for round := 1; round <= 10000; round++ {
		for _, monkey := range monkeys {
			throwItems(monkey, 2, p)
		}

		// just so we can see progress for a super long lasting run
		for _, monkey := range monkeys {
			fmt.Println(round, "monkey", monkey.id, "inspections", monkey.inspections, "items", monkey.items)
		}
		fmt.Println("")
	}

	sums := []int{}

	for _, monkey := range monkeys {
		fmt.Println("final result", "monkey", monkey.id, "inspections", monkey.inspections)
		sums = append(sums, monkey.inspections)
	}

	sort.Ints(sums)
	sum := sums[len(sums)-1] * sums[len(sums)-2]

	fmt.Println("sum", sum)
}

func createTestMonkeys() (*big.Int, []*Monkey) {
	zero := big.NewInt(0)
	b3 := big.NewInt(3)
	b6 := big.NewInt(6)
	b13 := big.NewInt(13)
	b17 := big.NewInt(17)
	b19 := big.NewInt(19)
	b23 := big.NewInt(23)
	p := big.NewInt(3 * 6 * 13 * 17 * 19 * 23)

	monkey0 := Monkey{
		id:          0,
		items:       []*big.Int{big.NewInt(79), big.NewInt(98)},
		inspections: 0,
		op: func(n *big.Int) {
			n.Mul(n, b19)
		},
		test: func(n *big.Int) bool {
			var res big.Int
			res.Mod(n, b23)
			return res.Cmp(zero) == 0
		},
	}
	monkey1 := Monkey{
		id:          1,
		items:       []*big.Int{big.NewInt(54), big.NewInt(65), big.NewInt(75), big.NewInt(74)},
		inspections: 0,
		op: func(n *big.Int) {
			n.Add(n, b6)
		},
		test: func(n *big.Int) bool {
			var res big.Int
			res.Mod(n, b19)
			return res.Cmp(zero) == 0
		},
	}
	monkey2 := Monkey{
		id:          2,
		items:       []*big.Int{big.NewInt(79), big.NewInt(60), big.NewInt(97)},
		inspections: 0,
		op: func(n *big.Int) {
			n.Mul(n, n)
		},
		test: func(n *big.Int) bool {
			var res big.Int
			res.Mod(n, b13)
			return res.Cmp(zero) == 0
		},
	}
	monkey3 := Monkey{
		id:          3,
		items:       []*big.Int{big.NewInt(74)},
		inspections: 0,
		op: func(n *big.Int) {
			n.Add(n, b3)
		},
		test: func(n *big.Int) bool {
			var res big.Int
			res.Mod(n, b17)
			return res.Cmp(zero) == 0
		},
	}

	monkey0.truemon = &monkey2
	monkey0.falsemon = &monkey3

	monkey1.truemon = &monkey2
	monkey1.falsemon = &monkey0

	monkey2.truemon = &monkey1
	monkey2.falsemon = &monkey3

	monkey3.truemon = &monkey0
	monkey3.falsemon = &monkey1

	return p, []*Monkey{
		&monkey0,
		&monkey1,
		&monkey2,
		&monkey3,
	}
}

func createMonkeys() (*big.Int, []*Monkey) {
	zero := big.NewInt(0)
	b2 := big.NewInt(2)
	b3 := big.NewInt(3)
	b4 := big.NewInt(4)
	b5 := big.NewInt(5)
	b6 := big.NewInt(6)
	b7 := big.NewInt(7)
	b8 := big.NewInt(8)
	b11 := big.NewInt(11)
	b13 := big.NewInt(13)
	b17 := big.NewInt(17)
	b19 := big.NewInt(19)
	p := big.NewInt(2 * 3 * 4 * 5 * 6 * 7 * 8 * 11 * 13 * 17 * 19)

	monkey0 := Monkey{
		id:          0,
		items:       []*big.Int{big.NewInt(66), big.NewInt(79)},
		inspections: 0,
		op: func(n *big.Int) {
			n.Mul(n, b11)
		},
		test: func(n *big.Int) bool {
			var res big.Int
			res.Mod(n, b7)
			return res.Cmp(zero) == 0
		},
	}
	monkey1 := Monkey{
		id:          1,
		items:       []*big.Int{big.NewInt(84), big.NewInt(94), big.NewInt(94), big.NewInt(81), big.NewInt(98), big.NewInt(75)},
		inspections: 0,
		op: func(n *big.Int) {
			n.Mul(n, b17)
		},
		test: func(n *big.Int) bool {
			var res big.Int
			res.Mod(n, b13)
			return res.Cmp(zero) == 0
		},
	}
	monkey2 := Monkey{
		id:          2,
		items:       []*big.Int{big.NewInt(85), big.NewInt(79), big.NewInt(59), big.NewInt(64), big.NewInt(79), big.NewInt(95), big.NewInt(67)},
		inspections: 0,
		op: func(n *big.Int) {
			n.Add(n, b8)
		},
		test: func(n *big.Int) bool {
			var res big.Int
			res.Mod(n, b5)
			return res.Cmp(zero) == 0
		},
	}
	monkey3 := Monkey{
		id:          3,
		items:       []*big.Int{big.NewInt(70)},
		inspections: 0,
		op: func(n *big.Int) {
			n.Add(n, b3)
		},
		test: func(n *big.Int) bool {
			var res big.Int
			res.Mod(n, b19)
			return res.Cmp(zero) == 0
		},
	}
	monkey4 := Monkey{
		id:          4,
		items:       []*big.Int{big.NewInt(57), big.NewInt(69), big.NewInt(78), big.NewInt(78)},
		inspections: 0,
		op: func(n *big.Int) {
			n.Add(n, b4)
		},
		test: func(n *big.Int) bool {
			var res big.Int
			res.Mod(n, b2)
			return res.Cmp(zero) == 0
		},
	}
	monkey5 := Monkey{
		id:          5,
		items:       []*big.Int{big.NewInt(65), big.NewInt(92), big.NewInt(60), big.NewInt(74), big.NewInt(72)},
		inspections: 0,
		op: func(n *big.Int) {
			n.Add(n, b7)
		},
		test: func(n *big.Int) bool {
			var res big.Int
			res.Mod(n, b11)
			return res.Cmp(zero) == 0
		},
	}
	monkey6 := Monkey{
		id:          6,
		items:       []*big.Int{big.NewInt(77), big.NewInt(91), big.NewInt(91)},
		inspections: 0,
		op: func(n *big.Int) {
			n.Mul(n, n)
		},
		test: func(n *big.Int) bool {
			var res big.Int
			res.Mod(n, b17)
			return res.Cmp(zero) == 0
		},
	}
	monkey7 := Monkey{
		id:          7,
		items:       []*big.Int{big.NewInt(76), big.NewInt(58), big.NewInt(57), big.NewInt(55), big.NewInt(67), big.NewInt(77), big.NewInt(54), big.NewInt(99)},
		inspections: 0,
		op: func(n *big.Int) {
			n.Add(n, b6)
		},
		test: func(n *big.Int) bool {
			var res big.Int
			res.Mod(n, b3)
			return res.Cmp(zero) == 0
		},
	}

	monkey0.truemon = &monkey6
	monkey0.falsemon = &monkey7

	monkey1.truemon = &monkey5
	monkey1.falsemon = &monkey2

	monkey2.truemon = &monkey4
	monkey2.falsemon = &monkey5

	monkey3.truemon = &monkey6
	monkey3.falsemon = &monkey0

	monkey4.truemon = &monkey0
	monkey4.falsemon = &monkey3

	monkey5.truemon = &monkey3
	monkey5.falsemon = &monkey4

	monkey6.truemon = &monkey1
	monkey6.falsemon = &monkey7

	monkey7.truemon = &monkey2
	monkey7.falsemon = &monkey1

	return p, []*Monkey{
		&monkey0,
		&monkey1,
		&monkey2,
		&monkey3,
		&monkey4,
		&monkey5,
		&monkey6,
		&monkey7,
	}
}

func throwItems(monkey *Monkey, part int, p *big.Int) {
	var item *big.Int

	for range monkey.items {
		item = monkey.items[0]

		monkey.op(item)

		if part == 1 {
			// divide worry by three
			item.Div(item, big.NewInt(3))
		} else {
			// modulus by minimum common product to keep number size down
			item.Mod(item, p)
		}

		if monkey.test(item) {
			monkey.truemon.items = append(monkey.truemon.items, item)
		} else {
			monkey.falsemon.items = append(monkey.falsemon.items, item)
		}

		monkey.items = monkey.items[1:]

		monkey.inspections++
	}
}

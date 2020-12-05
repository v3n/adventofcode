package main

import (
	"fmt"
	"strings"

	. "github.com/v3n/adventofcode/pkg"
)

var cases = Cases{
	Case{
		Data:     "BFFFBBFRRR",
		Expected: "70,7,567",
	},
	Case{
		Data:     "FFFBBBFRRR",
		Expected: "14,7,119",
	},
	Case{
		Data:     "BBFFBBFRLL",
		Expected: "102,4,820",
	},
}

var case1 = Cases{
	Case{
		Data:     input1,
		Expected: "892",
	},
}

var case2 = Cases{
	Case{
		Data:     input1,
		Expected: "625",
	},
}

type Location struct {
	Row, Column int
}

func (l *Location) ID() int {
	return (l.Row * 8) + l.Column
}

func (l *Location) String() string {
	return fmt.Sprintf("%d,%d,%d", l.Row, l.Column, l.ID())
}

func (l *Location) MarshalText() string {
	return fmt.Sprintf("%d,%d,%d", l.Row, l.Column, l.ID())
}

const (
	Front = 'F'
	Back  = 'B'
	Left  = 'L'
	Right = 'R'
)

const (
	MaxRowsBits    = 7
	MaxColumnsBits = 3
)

var BSPLookup = map[rune]int{
	Front: 0,
	Back:  1,
	Left:  0,
	Right: 1,
}

func solver(input string) interface{} {
	var (
		row uint8
		col uint8
	)

	for i, char := range input[0:7] {
		row = row | (uint8(BSPLookup[char]) << ((MaxRowsBits - 1) - i))
	}

	for i, char := range input[7:] {
		col = col | (uint8(BSPLookup[char]) << ((MaxColumnsBits - 1) - i))
	}

	loc := &Location{
		Row:    int(row),
		Column: int(col),
	}


	return loc
}

func makebsp(input string) interface{} {
	fragments := strings.Split(input, "\n")

	locs := make([]*Location, 128*8)

	for _, v := range fragments {
		loc := solver(v).(*Location)
		locs[loc.ID()] = loc
	}

	return locs
}

func IDFromCoords(r, c int) int {
	return (r * 8) + c
}

func part1(input string) interface{} {
	bsp := makebsp(input).([]*Location)

	for i := len(bsp) - 1;; i-- {
		if bsp[i] == nil {
			continue
		}

		return bsp[i].ID()
	}
}

// 145 == not right
func part2(input string) interface{} {
	bsp := makebsp(input).([]*Location)

	for i, v := range bsp {
		row := i / 8
		col := i % 8

		if v == nil {
			if !(bsp[i+1] != nil && bsp[i-1] != nil) {
				continue
			}

			loc := &Location{
				Row:    row,
				Column: col,
			}
			return loc.ID()
		}
	}

	return nil
}

func main() {
	Run(cases, solver)
	Run(case1, part1)
	Run(case2, part2)
}

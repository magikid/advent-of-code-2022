package adventofcode

import (
	"hilandchris.com/aoc2022/adventofcode/crates"
)

type Day5 struct {
	input *string
}

func (day5 *Day5) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day5.part1(), day5.part2()
}

func (d *Day5) part1() PuzzleAnswer {
	pa := PuzzleAnswer{part: 1}
	initialSupplyStacks, instructions := crates.ParseInput(*d.input)
	for _, instruction := range instructions {
		initialSupplyStacks.ExecuteCrateMover9000(instruction)
	}
	pa.answer = initialSupplyStacks.TopCrates()

	return pa
}

func (d *Day5) part2() PuzzleAnswer {
	pa := PuzzleAnswer{part: 2}
	initialSupplyStacks, instructions := crates.ParseInput(*d.input)
	for _, instruction := range instructions {
		initialSupplyStacks.ExecuteCrateMover9001(instruction)
	}
	pa.answer = initialSupplyStacks.TopCrates()

	return pa
}

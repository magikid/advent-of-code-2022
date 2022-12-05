package adventofcode

import (
	"fmt"
	"strings"

	"hilandchris.com/aoc2022/adventofcode/elfitems"
)

type Day3 struct {
	input *string
}

func (day3 *Day3) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day3.part1(), day3.part2()
}

func (d *Day3) part1() PuzzleAnswer {
	splitInput := strings.Split(*d.input, "\n")
	pa := PuzzleAnswer{part: 1}
	packs := make(elfitems.Packs, len(splitInput)-1)
	for i, line := range splitInput[0 : len(splitInput)-1] {
		p := elfitems.Pack{}
		p.AddToCompartments(line)
		packs[i] = p
	}
	pa.answer = fmt.Sprint(packs.CalculatePriorityOfDupes())

	return pa
}

func (d *Day3) part2() PuzzleAnswer {
	splitInput := strings.Split(*d.input, "\n")
	pa := PuzzleAnswer{part: 2}
	packs := make(elfitems.Packs, len(splitInput)-1)
	for i, line := range splitInput[0 : len(splitInput)-1] {
		p := elfitems.Pack{}
		p.AddToCompartments(line)
		packs[i] = p
	}

	pa.answer = fmt.Sprint(packs.SecurityBadgePriority())

	return pa
}

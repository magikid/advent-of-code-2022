package adventofcode

import (
	"fmt"
	"strings"

	"hilandchris.com/aoc2022/adventofcode/cleaners"
)

type Day4 struct {
	input *string
}

func (day4 *Day4) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day4.part1(), day4.part2()
}

func (d *Day4) part1() PuzzleAnswer {
	pa := PuzzleAnswer{part: 1}
	splitLines := strings.Split(*d.input, "\n")
	count := 0
	for _, line := range splitLines {
		if line == "" {
			continue
		}
		elves := cleaners.New(line)
		if elves[0].FullyContains(&elves[1]) || elves[1].FullyContains(&elves[0]) {
			count++
		}
	}
	pa.answer = fmt.Sprint(count)
	return pa
}

func (d *Day4) part2() PuzzleAnswer {
	pa := PuzzleAnswer{part: 2}
	splitLines := strings.Split(*d.input, "\n")
	count := 0
	for _, line := range splitLines {
		if line == "" {
			continue
		}
		elves := cleaners.New(line)
		if elves[0].Overlaps(&elves[1]) {
			count++
		}
	}
	pa.answer = fmt.Sprint(count)

	return pa
}

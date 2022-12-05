package adventofcode

import "errors"

type Day11 struct {
	input *string
}

func (day11 *Day11) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day11.part1(), day11.part2()
}

func (d *Day11) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day11) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

package adventofcode

import "errors"

type Day8 struct {
	input *string
}

func (day8 *Day8) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day8.part1(), day8.part2()
}

func (d *Day8) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day8) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

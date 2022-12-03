package adventofcode

import "errors"

type Day13 struct {
	input *string
}

func (day13 *Day13) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day13.part1(), day13.part2()
}

func (d *Day13) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day13) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

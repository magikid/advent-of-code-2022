package adventofcode

import "errors"

type Day16 struct {
	input *string
}

func (day16 *Day16) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day16.part1(), day16.part2()
}

func (d *Day16) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day16) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

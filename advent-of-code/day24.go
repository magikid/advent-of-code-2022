package adventofcode

import "errors"

type Day24 struct {
	input *string
}

func (day24 *Day24) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day24.part1(), day24.part2()
}

func (d *Day24) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day24) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

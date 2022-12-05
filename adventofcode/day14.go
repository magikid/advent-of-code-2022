package adventofcode

import "errors"

type Day14 struct {
	input *string
}

func (day14 *Day14) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day14.part1(), day14.part2()
}

func (d *Day14) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day14) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

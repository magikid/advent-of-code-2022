package adventofcode

import "errors"

type Day22 struct {
	input *string
}

func (day22 *Day22) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day22.part1(), day22.part2()
}

func (d *Day22) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day22) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

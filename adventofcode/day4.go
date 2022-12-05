package adventofcode

import "errors"

type Day4 struct {
	input *string
}

func (day4 *Day4) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day4.part1(), day4.part2()
}

func (d *Day4) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day4) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

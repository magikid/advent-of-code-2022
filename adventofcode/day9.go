package adventofcode

import "errors"

type Day9 struct {
	input *string
}

func (day9 *Day9) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day9.part1(), day9.part2()
}

func (d *Day9) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day9) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

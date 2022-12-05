package adventofcode

import "errors"

type Day18 struct {
	input *string
}

func (day18 *Day18) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day18.part1(), day18.part2()
}

func (d *Day18) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day18) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

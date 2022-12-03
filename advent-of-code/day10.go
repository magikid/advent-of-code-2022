package adventofcode

import "errors"

type Day10 struct {
	input *string
}

func (day10 *Day10) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day10.part1(), day10.part2()
}

func (d *Day10) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day10) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

package adventofcode

import "errors"

type Day20 struct {
	input *string
}

func (day20 *Day20) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day20.part1(), day20.part2()
}

func (d *Day20) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day20) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

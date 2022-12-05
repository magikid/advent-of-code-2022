package adventofcode

import "errors"

type Day5 struct {
	input *string
}

func (day5 *Day5) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day5.part1(), day5.part2()
}

func (d *Day5) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day5) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

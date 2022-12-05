package adventofcode

import "errors"

type Day7 struct {
	input *string
}

func (day7 *Day7) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day7.part1(), day7.part2()
}

func (d *Day7) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day7) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

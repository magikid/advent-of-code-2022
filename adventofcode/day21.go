package adventofcode

import "errors"

type Day21 struct {
	input *string
}

func (day21 *Day21) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day21.part1(), day21.part2()
}

func (d *Day21) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day21) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

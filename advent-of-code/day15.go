package adventofcode

import "errors"

type Day15 struct {
	input *string
}

func (day15 *Day15) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day15.part1(), day15.part2()
}

func (d *Day15) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day15) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

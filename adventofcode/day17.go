package adventofcode

import "errors"

type Day17 struct {
	input *string
}

func (day17 *Day17) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day17.part1(), day17.part2()
}

func (d *Day17) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day17) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

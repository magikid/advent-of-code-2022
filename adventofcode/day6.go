package adventofcode

import "errors"

type Day6 struct {
	input *string
}

func (day6 *Day6) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day6.part1(), day6.part2()
}

func (d *Day6) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day6) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

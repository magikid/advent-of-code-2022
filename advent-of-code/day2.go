package adventofcode

import "errors"

type Day2 struct {
	input *string
}

func (day2 *Day2) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day2.part1(), day2.part2()
}

func (d *Day2) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day2) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

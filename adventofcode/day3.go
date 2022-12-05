package adventofcode

import "errors"

type Day3 struct {
	input *string
}

func (day3 *Day3) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day3.part1(), day3.part2()
}

func (d *Day3) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day3) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

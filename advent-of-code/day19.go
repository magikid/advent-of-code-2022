package adventofcode

import "errors"

type Day19 struct {
	input *string
}

func (day19 *Day19) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day19.part1(), day19.part2()
}

func (d *Day19) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day19) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

package adventofcode

import "errors"

type Day25 struct {
	input *string
}

func (day25 *Day25) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day25.part1(), day25.part2()
}

func (d *Day25) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day25) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

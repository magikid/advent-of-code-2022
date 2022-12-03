package adventofcode

import "errors"

type Day23 struct {
	input *string
}

func (day23 *Day23) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day23.part1(), day23.part2()
}

func (d *Day23) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day23) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

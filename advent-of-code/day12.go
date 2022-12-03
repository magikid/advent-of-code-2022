package adventofcode

import "errors"

type Day12 struct {
	input *string
}

func (day12 *Day12) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day12.part1(), day12.part2()
}

func (d *Day12) part1() PuzzleAnswer {
	return PuzzleAnswer{part: 1, err: errors.New("not implemented")}
}

func (d *Day12) part2() PuzzleAnswer {
	return PuzzleAnswer{part: 2, err: errors.New("not implemented")}
}

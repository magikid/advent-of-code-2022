package adventofcode

import (
	"errors"
)

type Day1 struct {
	input *string
}

func (day1 *Day1) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day1.part1(), day1.part2()

}

func (day1 *Day1) part1() PuzzleAnswer {
	var ans PuzzleAnswer
	ans.part = 1
	ans.err = errors.New("not implemented")

	return ans
}

func (day1 *Day1) part2() PuzzleAnswer {
	var ans PuzzleAnswer
	ans.part = 2
	ans.err = errors.New("not implemented")

	return ans
}

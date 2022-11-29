package adventofcode

import (
	"fmt"
	"sync"
)

type Runnable interface {
	Run(*sync.WaitGroup) error
}

type Puzzle struct {
	day   int64
	input string
}

type PuzzleAnswer struct {
	part   int64
	answer string
	err    error
}

func (puzzle *Puzzle) Solve(wg *sync.WaitGroup) {
	defer wg.Done()
	ans1, ans2 := puzzle.SolveParts()

	ans1.PrintResults(puzzle.day)
	ans2.PrintResults(puzzle.day)
}

func (puzzle *Puzzle) SolveParts() (PuzzleAnswer, PuzzleAnswer) {
	switch puzzle.day {
	default:
		return PuzzleAnswer{}, PuzzleAnswer{}
	}
}

func (answer *PuzzleAnswer) PrintResults(day int64) {
	if answer.err != nil {
		fmt.Printf("❌ d%dp%d Error: %s\n", day, answer.part, answer.err.Error())
	}

	if answer.answer != "" {
		fmt.Printf("✅ d%dp%d Answer: %s\n", day, answer.part, answer.answer)
	}
}

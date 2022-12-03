package adventofcode

import (
	"log"
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
	case 1:
		puzzle := Day1{&puzzle.input}
		return puzzle.parts()
	case 2:
		puzzle := Day2{&puzzle.input}
		return puzzle.parts()
	case 3:
		puzzle := Day3{&puzzle.input}
		return puzzle.parts()
	case 4:
		puzzle := Day4{&puzzle.input}
		return puzzle.parts()
	case 5:
		puzzle := Day5{&puzzle.input}
		return puzzle.parts()
	case 6:
		puzzle := Day6{&puzzle.input}
		return puzzle.parts()
	case 7:
		puzzle := Day7{&puzzle.input}
		return puzzle.parts()
	case 8:
		puzzle := Day8{&puzzle.input}
		return puzzle.parts()
	case 9:
		puzzle := Day9{&puzzle.input}
		return puzzle.parts()
	case 10:
		puzzle := Day10{&puzzle.input}
		return puzzle.parts()
	case 11:
		puzzle := Day11{&puzzle.input}
		return puzzle.parts()
	case 12:
		puzzle := Day12{&puzzle.input}
		return puzzle.parts()
	case 13:
		puzzle := Day13{&puzzle.input}
		return puzzle.parts()
	case 14:
		puzzle := Day14{&puzzle.input}
		return puzzle.parts()
	case 15:
		puzzle := Day15{&puzzle.input}
		return puzzle.parts()
	case 16:
		puzzle := Day16{&puzzle.input}
		return puzzle.parts()
	case 17:
		puzzle := Day17{&puzzle.input}
		return puzzle.parts()
	case 18:
		puzzle := Day18{&puzzle.input}
		return puzzle.parts()
	case 19:
		puzzle := Day19{&puzzle.input}
		return puzzle.parts()
	case 20:
		puzzle := Day20{&puzzle.input}
		return puzzle.parts()
	case 21:
		puzzle := Day21{&puzzle.input}
		return puzzle.parts()
	case 22:
		puzzle := Day22{&puzzle.input}
		return puzzle.parts()
	case 23:
		puzzle := Day23{&puzzle.input}
		return puzzle.parts()
	case 24:
		puzzle := Day24{&puzzle.input}
		return puzzle.parts()
	case 25:
		puzzle := Day25{&puzzle.input}
		return puzzle.parts()
	default:
		return PuzzleAnswer{}, PuzzleAnswer{}
	}
}

func (answer *PuzzleAnswer) PrintResults(day int64) {
	if answer.err != nil {
		log.Printf("❌ d%dp%d Error: %s\n", day, answer.part, answer.err.Error())
	}

	if answer.answer != "" {
		log.Printf("✅ d%dp%d Answer: %s\n", day, answer.part, answer.answer)
	}
}

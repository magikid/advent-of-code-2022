package adventofcode

import (
	"fmt"
	"log"
	"strings"
)

type Day2 struct {
	input *string
}

func (day2 *Day2) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day2.part1(), day2.part2()
}

func (d *Day2) part1() PuzzleAnswer {
	lines := strings.Split(*d.input, "\n")
	strats := parseStrategies(lines, parseStrategyPart1)
	totalScore := 0
	for _, strat := range strats {
		totalScore += strat.scoreGame()
	}

	puzzAns := PuzzleAnswer{part: 1}
	puzzAns.answer = fmt.Sprint(totalScore)

	return puzzAns
}

func (d *Day2) part2() PuzzleAnswer {
	lines := strings.Split(*d.input, "\n")
	strats := parseStrategies(lines, parseStrategyPart2)
	totalScore := 0
	for _, strat := range strats {
		totalScore += strat.scoreGame()
	}

	puzzAns := PuzzleAnswer{part: 2}
	puzzAns.answer = fmt.Sprint(totalScore)

	return puzzAns
}

func parseStrategies(lines []string, parseStrategy func(string) RockPaperScissorsGame) []RockPaperScissorsGame {
	games := make([]RockPaperScissorsGame, len(lines)-1)
	for i, line := range lines {
		if line == "" {
			log.Printf("Skipping line number: %d", i)
			continue
		}
		games[i] = parseStrategy(line)
	}

	return games
}

func parseStrategyPart1(line string) RockPaperScissorsGame {
	moves := strings.Split(line, " ")
	opponentRawMove := moves[0]
	opponentMove := Rock
	switch opponentRawMove {
	case "A":
		opponentMove = Rock
	case "B":
		opponentMove = Paper
	case "C":
		opponentMove = Scissors
	}
	myRawMove := moves[1]
	myMove := Rock
	switch myRawMove {
	case "X":
		myMove = Rock
	case "Y":
		myMove = Paper
	case "Z":
		myMove = Scissors
	}

	return RockPaperScissorsGame{myMove: myMove, opponentMove: opponentMove}
}

func parseStrategyPart2(line string) RockPaperScissorsGame {
	moves := strings.Split(line, " ")
	opponentRawMove := moves[0]
	gameOutcome := moves[1]
	opponentMove := Rock
	myMove := Rock
	switch opponentRawMove {
	case "A":
		opponentMove = Rock
		switch gameOutcome {
		case "X":
			myMove = Scissors
		case "Y":
			myMove = Rock
		case "Z":
			myMove = Paper
		}
	case "B":
		opponentMove = Paper
		switch gameOutcome {
		case "X":
			myMove = Rock
		case "Y":
			myMove = Paper
		case "Z":
			myMove = Scissors
		}
	case "C":
		opponentMove = Scissors
		switch gameOutcome {
		case "X":
			myMove = Paper
		case "Y":
			myMove = Scissors
		case "Z":
			myMove = Rock
		}
	}

	return RockPaperScissorsGame{myMove: myMove, opponentMove: opponentMove}
}

type Move int
type Outcome int

const (
	Rock Move = iota
	Paper
	Scissors
)

const (
	Win Outcome = iota
	Lose
	Tie
)

type RockPaperScissorsGame struct {
	opponentMove Move
	myMove       Move
}

func (game *RockPaperScissorsGame) scoreGame() int {
	return game.calculateMoveScore() + game.calculateOutcomeScore()
}

func (game *RockPaperScissorsGame) calculateMoveScore() int {
	switch game.myMove {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	default:
		return 0
	}
}

func (game *RockPaperScissorsGame) calculateOutcomeScore() int {
	outcome := game.doIWin()
	switch outcome {
	case Win:
		return 6
	case Tie:
		return 3
	default:
		return 0
	}
}

func (game *RockPaperScissorsGame) doIWin() Outcome {
	switch game.myMove {
	case Rock:
		switch game.opponentMove {
		case Rock:
			return Tie
		case Scissors:
			return Win
		case Paper:
			return Lose
		}
	case Paper:
		switch game.opponentMove {
		case Rock:
			return Win
		case Scissors:
			return Lose
		case Paper:
			return Tie
		}
	case Scissors:
		switch game.opponentMove {
		case Rock:
			return Lose
		case Scissors:
			return Tie
		case Paper:
			return Win
		}
	}
	return Lose
}

package adventofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionality(t *testing.T) {
	game := RockPaperScissorsGame{myMove: Paper, opponentMove: Rock}
	assert.Equal(t, 8, game.scoreGame())

	game = RockPaperScissorsGame{myMove: Rock, opponentMove: Paper}
	assert.Equal(t, 1, game.scoreGame())

	game = RockPaperScissorsGame{myMove: Scissors, opponentMove: Scissors}
	assert.Equal(t, 6, game.scoreGame())
}

func TestParseMove(t *testing.T) {
	line := "A Y"
	assert.Equal(t, RockPaperScissorsGame{myMove: Paper, opponentMove: Rock}, parseStrategyPart1(line))

	line = "B X"
	assert.Equal(t, RockPaperScissorsGame{opponentMove: Paper, myMove: Rock}, parseStrategyPart1(line))

	line = "C Z"
	assert.Equal(t, RockPaperScissorsGame{opponentMove: Scissors, myMove: Scissors}, parseStrategyPart1(line))

}

func TestStratScore(t *testing.T) {
	lines := []string{"A Y", "B X", "C Z", ""}
	games := parseStrategies(lines, parseStrategyPart1)
	total := 0
	for _, game := range games {
		total += game.scoreGame()
	}
	assert.Equal(t, 15, total)
}

func TestNewStratScore(t *testing.T) {
	lines := []string{"A Y", "B X", "C Z", ""}
	games := parseStrategies(lines, parseStrategyPart2)
	total := 0
	for _, game := range games {
		total += game.scoreGame()
	}
	assert.Equal(t, 12, total)
}

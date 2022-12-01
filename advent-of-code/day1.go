package adventofcode

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Day1 struct {
	input *string
}

type Pack struct {
	snackCalories []int
	totalCalories int
}

func (p *Pack) addFood(calories int) {
	p.snackCalories = append(p.snackCalories, calories)
}

func (p *Pack) calculateTotalCalories() int {
	if p.totalCalories != 0 {
		return p.totalCalories
	}

	totalCalories := 0
	for _, calorieCount := range p.snackCalories {
		totalCalories += calorieCount
	}
	p.totalCalories = totalCalories

	return p.totalCalories
}

func (day1 *Day1) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day1.part1(), day1.part2()

}

func (day1 *Day1) part1() PuzzleAnswer {
	var ans PuzzleAnswer
	ans.part = 1
	maxPack := 0
	pack := Pack{}
	lines := strings.Split(*day1.input, "\n")

	for _, line := range lines {
		if len(line) == 0 {
			if pack.calculateTotalCalories() > maxPack {
				maxPack = pack.totalCalories
			}
			pack = Pack{}
		}
		calories, _ := strconv.Atoi(line)
		pack.addFood(calories)
	}
	ans.answer = fmt.Sprint(maxPack)

	return ans
}

func (day1 *Day1) part2() PuzzleAnswer {
	var ans PuzzleAnswer
	ans.part = 2
	ans.err = errors.New("not implemented")

	return ans
}

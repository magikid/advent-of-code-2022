package adventofcode

import (
	"fmt"
	"sort"
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
	currentPack := Pack{}
	packs := []Pack{currentPack}
	lines := strings.Split(*day1.input, "\n")

	for _, line := range lines {
		if len(line) == 0 {
			currentPack.calculateTotalCalories()
			packs = append(packs, currentPack)
			currentPack = Pack{}
		}
		calories, _ := strconv.Atoi(line)
		currentPack.addFood(calories)
	}

	sort.Slice(packs, func(p, q int) bool {
		return packs[p].totalCalories > packs[q].totalCalories
	})

	ans.answer = fmt.Sprint(packs[0].totalCalories)

	return ans
}

func (day1 *Day1) part2() PuzzleAnswer {
	var ans PuzzleAnswer
	ans.part = 2
	currentPack := Pack{}
	packs := []Pack{currentPack}
	lines := strings.Split(*day1.input, "\n")

	for _, line := range lines {
		if len(line) == 0 {
			currentPack.calculateTotalCalories()
			packs = append(packs, currentPack)
			currentPack = Pack{}
		}
		calories, _ := strconv.Atoi(line)
		currentPack.addFood(calories)
	}

	sort.Slice(packs, func(p, q int) bool {
		return packs[p].totalCalories > packs[q].totalCalories
	})

	ans.answer = fmt.Sprintf("Calories: %d + %d + %d = %d", packs[0].totalCalories, packs[1].totalCalories, packs[2].totalCalories, packs[0].totalCalories+packs[1].totalCalories+packs[2].totalCalories)

	return ans
}

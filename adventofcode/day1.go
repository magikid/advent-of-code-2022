package adventofcode

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"hilandchris.com/aoc2022/adventofcode/elfitems"
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
	currentPack := elfitems.Pack{}
	packs := elfitems.Packs{currentPack}
	lines := strings.Split(*day1.input, "\n")

	for _, line := range lines {
		if len(line) == 0 {
			currentPack.CalculateTotalCalories()
			packs = append(packs, currentPack)
			currentPack = elfitems.Pack{}
		}
		calories, _ := strconv.Atoi(line)
		currentPack.AddCalories(calories)
	}

	sort.Sort(packs)

	ans.answer = fmt.Sprint(packs[0].CalculateTotalCalories())

	return ans
}

func (day1 *Day1) part2() PuzzleAnswer {
	var ans PuzzleAnswer
	ans.part = 2
	currentPack := elfitems.Pack{}
	packs := elfitems.Packs{currentPack}
	lines := strings.Split(*day1.input, "\n")

	for _, line := range lines {
		if len(line) == 0 {
			currentPack.CalculateTotalCalories()
			packs = append(packs, currentPack)
			currentPack = elfitems.Pack{}
		}
		calories, _ := strconv.Atoi(line)
		currentPack.AddCalories(calories)
	}

	sort.Slice(packs, func(p, q int) bool {
		return packs[p].CalculateTotalCalories() > packs[q].CalculateTotalCalories()
	})

	ans.answer = fmt.Sprintf("Calories: %d + %d + %d = %d", packs[0].CalculateTotalCalories(), packs[1].CalculateTotalCalories(), packs[2].CalculateTotalCalories(), packs[0].CalculateTotalCalories()+packs[1].CalculateTotalCalories()+packs[2].CalculateTotalCalories())

	return ans
}

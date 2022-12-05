package elfitems

import (
	"unicode"
)

func (p *Pack) AddCalories(calories int) {
	p.snackCalories = append(p.snackCalories, calories)
}

func (p *Pack) CalculateTotalCalories() int {
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

func (p *Pack) AddToCompartments(items string) {
	p.packItems = convertPackItems(items)
}

func (p *Pack) firstCompartment() PackItems {
	return p.packItems[0:(len(p.packItems) / 2)]
}

func (p *Pack) secondCompartment() PackItems {
	return p.packItems[(len(p.packItems) / 2):]
}

func (p *Pack) findDuplicateItemTypes() PackItem {
	for _, first_item_type := range p.firstCompartment() {
		for _, second_item_type := range p.secondCompartment() {
			if first_item_type == second_item_type {
				return PackItem(second_item_type)
			}
		}
	}
	return PackItem("")
}

func (p *Pack) calculateDuplicateItemPriority() int {
	firstRuneInPackItem := []rune(p.findDuplicateItemTypes())[0]

	var asciiCodeOfDuplicate rune
	if !unicode.IsUpper(firstRuneInPackItem) {
		asciiCodeOfDuplicate = firstRuneInPackItem - 96
	} else {
		asciiCodeOfDuplicate = firstRuneInPackItem - 38
	}

	return int(asciiCodeOfDuplicate)
}

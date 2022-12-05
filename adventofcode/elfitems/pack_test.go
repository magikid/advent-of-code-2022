package elfitems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElfPack(t *testing.T) {
	p := Pack{snackCalories: []int{1000, 2000, 3000}}
	assert.Equal(t, 6000, p.CalculateTotalCalories())
	assert.Equal(t, 6000, p.totalCalories)

	p = Pack{snackCalories: []int{4000}}
	assert.Equal(t, 4000, p.CalculateTotalCalories())

	p = Pack{snackCalories: []int{7000, 8000, 9000}}
	assert.Equal(t, 24000, p.CalculateTotalCalories())

	p = Pack{snackCalories: []int{10000}}
	assert.Equal(t, 10000, p.CalculateTotalCalories())
}

func TestCompartment(t *testing.T) {
	p := Pack{}
	p.AddToCompartments("vJrwpWtwJgWrhcsFMMfFFhFp")
	assert.Equal(t, convertPackItems("vJrwpWtwJgWr"), p.firstCompartment())
	assert.Equal(t, convertPackItems("hcsFMMfFFhFp"), p.secondCompartment())
	assert.Equal(t, PackItem("p"), p.findDuplicateItemTypes())

	p.AddToCompartments("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL")
	assert.Equal(t, convertPackItems("jqHRNqRjqzjGDLGL"), p.firstCompartment())
	assert.Equal(t, convertPackItems("rsFMfFZSrLrFZsSL"), p.secondCompartment())
	assert.Equal(t, PackItem("L"), p.findDuplicateItemTypes())

	p.AddToCompartments("PmmdzqPrVvPwwTWBwg")
	assert.Equal(t, convertPackItems("PmmdzqPrV"), p.firstCompartment())
	assert.Equal(t, convertPackItems("vPwwTWBwg"), p.secondCompartment())
	assert.Equal(t, PackItem("P"), p.findDuplicateItemTypes())

	p.AddToCompartments("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn")
	assert.Equal(t, PackItem("v"), p.findDuplicateItemTypes())

	p.AddToCompartments("ttgJtRGJQctTZtZT")
	assert.Equal(t, PackItem("t"), p.findDuplicateItemTypes())

	p.AddToCompartments("CrZsJsPPZsGzwwsLwLmpwMDw")
	assert.Equal(t, PackItem("s"), p.findDuplicateItemTypes())
}

func TestItemPriority(t *testing.T) {
	p := Pack{}
	p.AddToCompartments("vJrwpWtwJgWrhcsFMMfFFhFp")
	assert.Equal(t, 16, p.calculateDuplicateItemPriority())

	p.AddToCompartments("aa")
	assert.Equal(t, 1, p.calculateDuplicateItemPriority())

	p.AddToCompartments("zz")
	assert.Equal(t, 26, p.calculateDuplicateItemPriority())

	p.AddToCompartments("AA")
	assert.Equal(t, 27, p.calculateDuplicateItemPriority())

	p.AddToCompartments("ZZ")
	assert.Equal(t, 52, p.calculateDuplicateItemPriority())
}

func TestSecurityBadge(t *testing.T) {
	p1 := Pack{}
	p1.AddToCompartments("vJrwpWtwJgWrhcsFMMfFFhFp")

	p2 := Pack{}
	p2.AddToCompartments("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL")

	p3 := Pack{}
	p3.AddToCompartments("PmmdzqPrVvPwwTWBwg")

	packs := Packs{p1, p2, p3}
	assert.Equal(t, 18, packs.SecurityBadgePriority())

	p1.AddToCompartments("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn")
	p2.AddToCompartments("ttgJtRGJQctTZtZT")
	p3.AddToCompartments("CrZsJsPPZsGzwwsLwLmpwMDw")

	packs = Packs{p1, p2, p3}
	assert.Equal(t, 52, packs.SecurityBadgePriority())
}

func TestByThrees(t *testing.T) {
	p1 := Pack{}
	p1.AddToCompartments("vJrwpWtwJgWrhcsFMMfFFhFp")

	p2 := Pack{}
	p2.AddToCompartments("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL")

	p3 := Pack{}
	p3.AddToCompartments("PmmdzqPrVvPwwTWBwg")

	packs := Packs{p1, p2, p3, p3, p2, p1}

	groupedPacks := packs.groupByThree()

	assert.Equal(t, 3, len(groupedPacks[0]))
	assert.Equal(t, 3, len(groupedPacks[1]))
	assert.Equal(t, 2, len(groupedPacks))
	assert.Equal(t, Packs{p1, p2, p3}, groupedPacks[0])
	assert.Equal(t, Packs{p3, p2, p1}, groupedPacks[1])
}

package elfitems

import (
	"sort"
)

func (p Packs) Len() int {
	return len(p)
}

func (p Packs) Less(i, j int) bool {
	return p[i].totalCalories > p[j].totalCalories
}

func (p Packs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (pi PackItem) String() string {
	return string(pi)
}

func (pis PackItems) String() string {
	output := ""
	for _, item := range pis {
		output += string(item)
	}

	return output
}

func (packs Packs) CalculatePriorityOfDupes() int {
	total := 0
	for _, pack := range packs {
		total += pack.calculateDuplicateItemPriority()
	}

	return total
}

func (packs Packs) SecurityBadgePriority() int {
	total := 0
	groupPriority := 0
	for _, group := range packs.groupByThree() {
		sort.Sort(ByLenPackItems(group))

		for _, letter1 := range group[0].packItems {
			for _, letter2 := range group[1].packItems {
				for _, letter3 := range group[2].packItems {
					if letter1 == letter2 && letter1 == letter3 {
						groupPriority = PackItemPriority([]rune(PackItem(letter3))[0])
						break
					}
				}
			}
		}

		total += groupPriority
		groupPriority = 0
	}

	return total
}

func (packs Packs) groupByThree() []Packs {
	output := make([]Packs, len(packs)/3)
	group := Packs{}
	groupCount := 0
	for i := 1; i <= len(packs); i++ {
		group = append(group, packs[i-1])
		if i%3 == 0 {
			output[groupCount] = group
			groupCount++
			group = Packs{}
		}
	}

	return output
}

type ByLenPackItems Packs

func (a ByLenPackItems) Len() int           { return len(a) }
func (a ByLenPackItems) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLenPackItems) Less(i, j int) bool { return len(a[i].packItems) > len(a[j].packItems) }

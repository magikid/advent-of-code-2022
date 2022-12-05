package elfitems

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

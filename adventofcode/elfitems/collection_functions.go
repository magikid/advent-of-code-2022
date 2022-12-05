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

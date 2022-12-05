package elfitems

func (p PackItems) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p PackItems) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PackItems) Len() int {
	return len(p)
}

func convertPackItems(line string) PackItems {
	items := make(PackItems, len(line))

	for i, letter := range line {
		items[i] = PackItem(string(letter))
	}

	return items
}

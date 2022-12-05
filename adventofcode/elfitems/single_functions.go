package elfitems

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

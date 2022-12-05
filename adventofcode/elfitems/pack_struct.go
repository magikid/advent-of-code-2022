package elfitems

type Packs []Pack

type Pack struct {
	snackCalories []int
	totalCalories int
	packItems     PackItems
}

type PackItems []PackItem
type PackItem string
type PackItemPriority int

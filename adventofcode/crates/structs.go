package crates

type Crate string
type Stack []Crate
type SupplyStacks []Stack
type craneInstruction struct {
	numberOfCrates int
	fromStack      int
	toStack        int
}

package crates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	rawInst := "move 1 from 2 to 1"
	instr := parseInstructions(rawInst)
	assert.Equal(t, 1, instr.numberOfCrates)
	assert.Equal(t, 2, instr.fromStack)
	assert.Equal(t, 1, instr.toStack)
}

func TestExecute(t *testing.T) {
	stacks := SupplyStacks{Stack{Crate("A"), Crate("B")}, Stack{Crate("C")}}
	instr := craneInstruction{numberOfCrates: 1, fromStack: 0, toStack: 1}
	newStacks := stacks.execute(instr)

	assert.Equal(t, SupplyStacks{Stack{Crate("A")}, Stack{Crate("C"), Crate("B")}}, newStacks)
}

package crates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	rawInst := "move 1 from 2 to 1"
	instr := parseInstruction(rawInst)
	assert.Equal(t, 1, instr.numberOfCrates)
	assert.Equal(t, 2, instr.fromStack)
	assert.Equal(t, 1, instr.toStack)
}

func TestExecute9000(t *testing.T) {
	stacks := SupplyStacks{Stack{Crate("A"), Crate("B")}, Stack{Crate("C")}}
	instr := craneInstruction{numberOfCrates: 1, fromStack: 1, toStack: 2}
	newStacks := stacks.ExecuteCrateMover9000(instr)

	assert.Equal(t, SupplyStacks{Stack{Crate("A")}, Stack{Crate("C"), Crate("B")}}, newStacks)
}

func TestExecute9001(t *testing.T) {
	stacks := SupplyStacks{Stack{Crate("A"), Crate("B"), Crate("C"), Crate("D")}, Stack{Crate("E")}}
	instr := craneInstruction{numberOfCrates: 2, fromStack: 1, toStack: 2}
	newStacks := stacks.ExecuteCrateMover9001(instr)

	assert.Equal(t, SupplyStacks{Stack{Crate("A"), Crate("B")}, Stack{Crate("E"), Crate("C"), Crate("D")}}, newStacks)
}

func TestParseSupplyStack(t *testing.T) {
	input := "[A] [B] [C] [D] [E] [F] [G] [H] [I]\n[N] [M] [F] [D] [R] [C] [W] [T] [M]\n 1   2   3   4   5   6   7   8   9 "
	stack := parseSupplyStack(input)

	assert.Equal(t, Crate("N"), stack[0][0])
}

func TestTop9000Crates(t *testing.T) {
	input := `   	    [M]     [B]             [N]
[T]     [H]     [V] [Q]         [H]
[Q]     [N]     [H] [W] [T]     [Q]
[V]     [P] [F] [Q] [P] [C]     [R]
[C]     [D] [T] [N] [N] [L] [S] [J]
[D] [V] [W] [R] [M] [G] [R] [N] [D]
[S] [F] [Q] [Q] [F] [F] [F] [Z] [S]
[N] [M] [F] [D] [R] [C] [W] [T] [M]
 1   2   3   4   5   6   7   8   9

move 1 from 8 to 7`
	stack, instructions := ParseInput(input)
	assert.Equal(t, stack[0], Stack{Crate("N"), Crate("S"), Crate("D"), Crate("C"), Crate("V"), Crate("Q"), Crate("T")})
	assert.Equal(t, "TVMFBQTSN", stack.TopCrates())
	for _, i := range instructions {
		stack = stack.ExecuteCrateMover9000(i)
	}
	assert.Equal(t, "TVMFBQSNN", stack.TopCrates())
}

func TestPopN(t *testing.T) {
	stack := Stack{Crate("A"), Crate("B"), Crate("C")}
	crates, stack := stack.popN(2)
	assert.Equal(t, Stack{Crate("A")}, stack)
	assert.Equal(t, []Crate{Crate("B"), Crate("C")}, crates)
}

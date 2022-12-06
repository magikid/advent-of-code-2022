package crates

import (
	"regexp"
	"strconv"
)

func parseInstructions(line string) craneInstruction {
	pattern := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	matches := pattern.FindStringSubmatch(line)
	numberOfCrates, _ := strconv.Atoi(matches[1])
	fromStack, _ := strconv.Atoi(matches[2])
	toStack, _ := strconv.Atoi(matches[3])

	return craneInstruction{numberOfCrates: numberOfCrates, fromStack: fromStack, toStack: toStack}
}

func (stack Stack) pop() (Crate, Stack) {
	i := len(stack) - 1
	crate := stack[i]
	stack = append(stack[:i], stack[i+1:]...)

	return crate, stack
}

func (stack Stack) push(crate Crate) Stack {
	stack = append(stack, crate)
	return stack
}

func (stacks SupplyStacks) execute(instr craneInstruction) SupplyStacks {
	for i := 0; i < instr.numberOfCrates; i++ {
		crate, newStack := stacks[instr.fromStack].pop()
		stacks[instr.fromStack] = newStack
		stacks[instr.toStack] = stacks[instr.toStack].push(crate)
	}

	return stacks
}

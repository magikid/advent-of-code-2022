package crates

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	FirstColumn   int = 1
	SecondColumn  int = 5
	ThirdColumn   int = 9
	FourthColumn  int = 13
	FifthColumn   int = 17
	SixthColumn   int = 21
	SeventhColumn int = 25
	EigthColumn   int = 29
	NinthColumn   int = 33
)

func ParseInput(input string) (SupplyStacks, []craneInstruction) {
	splitInput := strings.Split(input, "\n\n")
	supplyStack := parseSupplyStack(splitInput[0])
	instructions := parseInstructions(strings.Split(splitInput[1], "\n"))

	return supplyStack, instructions
}

func parseInstructions(lines []string) []craneInstruction {
	instructions := make([]craneInstruction, len(lines))
	for i, line := range lines {
		if line == "" {
			continue
		}
		instructions[i] = parseInstruction(line)
	}

	return instructions
}

func parseSupplyStack(input string) SupplyStacks {
	rows := strings.Split(input, "\n")
	supplyStack := make(SupplyStacks, 9)

	for i := len(rows) - 2; i >= 0; i-- {
		firstColumn := string(rows[i][FirstColumn])
		if firstColumn != " " {
			supplyStack[0] = supplyStack[0].push(Crate(firstColumn))
		}

		secondColumn := string(rows[i][SecondColumn])
		if secondColumn != " " {
			supplyStack[1] = supplyStack[1].push(Crate(secondColumn))
		}

		thirdColumn := string(rows[i][ThirdColumn])
		if thirdColumn != " " {
			supplyStack[2] = supplyStack[2].push(Crate(thirdColumn))
		}

		fourthColumn := string(rows[i][FourthColumn])
		if fourthColumn != " " {
			supplyStack[3] = supplyStack[3].push(Crate(fourthColumn))
		}

		fifthColumn := string(rows[i][FifthColumn])
		if fifthColumn != " " {
			supplyStack[4] = supplyStack[4].push(Crate(fifthColumn))
		}

		sixthColumn := string(rows[i][SixthColumn])
		if sixthColumn != " " {
			supplyStack[5] = supplyStack[5].push(Crate(sixthColumn))
		}

		seventhColumn := string(rows[i][SeventhColumn])
		if seventhColumn != " " {
			supplyStack[6] = supplyStack[6].push(Crate(seventhColumn))
		}

		eightColumn := string(rows[i][EigthColumn])
		if eightColumn != " " {
			supplyStack[7] = supplyStack[7].push(Crate(eightColumn))
		}

		ninthColumn := string(rows[i][NinthColumn])
		if ninthColumn != " " {
			supplyStack[8] = supplyStack[8].push(Crate(ninthColumn))
		}
	}

	return supplyStack
}

func parseInstruction(line string) craneInstruction {
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

func (stack Stack) popN(n int) ([]Crate, Stack) {
	crates := make([]Crate, n)
	crate := Crate("")
	for i := n - 1; i >= 0; i-- {
		crate, stack = stack.pop()
		crates[i] = crate
	}

	return crates, stack
}

func reverse(crates []Crate) []Crate {
	output := make([]Crate, len(crates))

	for i, crate := range crates {
		j := len(crates) - i - 1
		output[j] = crate
	}

	return output
}

func (stack Stack) push(crate Crate) Stack {
	stack = append(stack, crate)
	return stack
}

func (stack Stack) pushMany(crate []Crate) Stack {
	stack = append(stack, crate...)
	return stack
}

func (stacks SupplyStacks) ExecuteCrateMover9000(instr craneInstruction) SupplyStacks {
	for i := 0; i < instr.numberOfCrates; i++ {
		crate, newStack := stacks[instr.fromStack-1].pop()
		stacks[instr.fromStack-1] = newStack
		stacks[instr.toStack-1] = stacks[instr.toStack-1].push(crate)
	}

	return stacks
}

func (stacks SupplyStacks) ExecuteCrateMover9001(instr craneInstruction) SupplyStacks {
	if instr.fromStack == 0 || instr.toStack == 0 || instr.numberOfCrates == 0 {
		return stacks
	}

	crates, newStack := stacks[instr.fromStack-1].popN(instr.numberOfCrates)
	stacks[instr.fromStack-1] = newStack
	stacks[instr.toStack-1] = stacks[instr.toStack-1].pushMany(crates)

	return stacks
}

func (stacks SupplyStacks) TopCrates() string {
	topCrates := make([]string, len(stacks))

	squashedStack := Stack{}
	for i, stack := range stacks {
		for _, crate := range stack {
			if crate == Crate(" ") {
				continue
			}
			squashedStack = append(squashedStack, crate)
		}

		topCrates[i] = string(squashedStack[len(squashedStack)-1])
		squashedStack = Stack{}
	}

	return strings.Join(topCrates, "")
}

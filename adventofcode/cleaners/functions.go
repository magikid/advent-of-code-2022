package cleaners

import (
	"strconv"
	"strings"
)

func New(line string) []Elf {
	elves := make([]Elf, 2)
	assignments := strings.Split(line, ",")
	firstElfSections := strings.Split(assignments[0], "-")
	start, _ := strconv.Atoi(firstElfSections[0])
	end, _ := strconv.Atoi(firstElfSections[1])
	elves[0] = Elf{assignmentStart: start, assignmentEnd: end}

	secondElfSections := strings.Split(assignments[1], "-")
	start, _ = strconv.Atoi(secondElfSections[0])
	end, _ = strconv.Atoi(secondElfSections[1])
	elves[1] = Elf{assignmentStart: start, assignmentEnd: end}

	return elves
}

func (elf1 *Elf) FullyContains(elf2 *Elf) bool {
	return elf1.assignmentStart <= elf2.assignmentStart && elf1.assignmentEnd >= elf2.assignmentEnd
}

func (elf1 *Elf) Overlaps(elf2 *Elf) bool {
	for i := elf1.assignmentStart; i <= elf1.assignmentEnd; i++ {
		for j := elf2.assignmentStart; j <= elf2.assignmentEnd; j++ {
			if i == j {
				return true
			}
		}
	}

	return false
}

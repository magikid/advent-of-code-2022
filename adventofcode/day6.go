package adventofcode

import (
	"fmt"

	"hilandchris.com/aoc2022/adventofcode/comms"
)

type Day6 struct {
	input *string
}

func (day6 *Day6) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day6.part1(), day6.part2()
}

func (d *Day6) part1() PuzzleAnswer {
	pa := PuzzleAnswer{part: 1}
	datastream := comms.Datastream{DatastreamBuffer: *d.input}
	pa.answer = fmt.Sprint(datastream.StartOfPacketMarkerLocation())

	return pa
}

func (d *Day6) part2() PuzzleAnswer {
	pa := PuzzleAnswer{part: 2}
	datastream := comms.Datastream{DatastreamBuffer: *d.input}
	pa.answer = fmt.Sprint(datastream.StartOfMessageMarkerLocation())

	return pa
}

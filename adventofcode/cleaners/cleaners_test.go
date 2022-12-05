package cleaners

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	line := "2-4,6-8"
	elves := New(line)
	assert.Equal(t, 2, elves[0].assignmentStart)
	assert.Equal(t, 4, elves[0].assignmentEnd)
	assert.Equal(t, 6, elves[1].assignmentStart)
	assert.Equal(t, 8, elves[1].assignmentEnd)
}

func TestContains(t *testing.T) {
	line := "2-8,3-7"
	elves := New(line)
	assert.True(t, elves[0].FullyContains(&elves[1]))

	line = "6-6,4-6"
	elves = New(line)
	assert.True(t, elves[1].FullyContains(&elves[0]))
}

func TestOverlaps(t *testing.T) {
	line := "5-7,6-9"
	elves := New(line)
	assert.True(t, elves[0].Overlaps(&elves[1]))

	line = "2-8,3-7"
	elves = New(line)
	assert.True(t, elves[0].Overlaps(&elves[1]))

	line = "6-6,4-6"
	elves = New(line)
	assert.True(t, elves[0].Overlaps(&elves[1]))

	line = "2-6,4-8"
	elves = New(line)
	assert.True(t, elves[0].Overlaps(&elves[1]))
}

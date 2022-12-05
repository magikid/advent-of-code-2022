package elfitems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElfPack(t *testing.T) {
	p := Pack{snackCalories: []int{1000, 2000, 3000}}
	assert.Equal(t, 6000, p.CalculateTotalCalories())
	assert.Equal(t, 6000, p.totalCalories)

	p = Pack{snackCalories: []int{4000}}
	assert.Equal(t, 4000, p.CalculateTotalCalories())

	p = Pack{snackCalories: []int{7000, 8000, 9000}}
	assert.Equal(t, 24000, p.CalculateTotalCalories())

	p = Pack{snackCalories: []int{10000}}
	assert.Equal(t, 10000, p.CalculateTotalCalories())
}

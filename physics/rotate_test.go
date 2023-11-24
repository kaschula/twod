package physics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rotateValue(t *testing.T) {
	assert.Equal(t, 1, CycleValue(360, 361))
	assert.Equal(t, 10, CycleValue(360, 370))
	assert.Equal(t, 10, CycleValue(360, -350))
}

func Test_indexCycler(t *testing.T) {

	assert.Equal(t, 0, indexCycler(3, 0))
	assert.Equal(t, 1, indexCycler(3, 1))
	assert.Equal(t, 2, indexCycler(3, 2))
	assert.Equal(t, 0, indexCycler(3, 3))
	assert.Equal(t, 1, indexCycler(3, 4))
	assert.Equal(t, 2, indexCycler(3, 5))
}

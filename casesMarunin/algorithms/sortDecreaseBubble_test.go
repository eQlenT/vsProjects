package algorithms

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestSortDecreaseBubble(t *testing.T) {
	expected := []int{1000, 679, 100, 52, 10, 0, -8, -56, -100, -1000, -2000}
	actual := SortDecreaseBubble([]int{10, -1000, -56, 100, -100, -2000, -8, 0, 52, 679, 1000})
	assert.Equal(t, expected, actual)
}
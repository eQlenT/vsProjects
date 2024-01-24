package algorithms

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
  
func TestSortDecreaseSlices(t *testing.T) {
	expected := []int{1000, 679, 100, 52, 10, 0, -8, -56, -100, -1000, -2000}
	actual := SortDecreaseSlices([]int{10, -1000, -56, 100, -100, -2000, -8, 0, 52, 679, 1000})
	assert.Equal(t, expected, actual)
}

func TestSortDecreaseSort(t *testing.T) {
	expected := []int{1000, 679, 100, 52, 10, 0, -8, -56, -100, -1000, -2000}
	actual := SortDecreaseSort([]int{10, -1000, -56, 100, -100, -2000, -8, 0, 52, 679, 1000})
	assert.Equal(t, expected, actual)
}
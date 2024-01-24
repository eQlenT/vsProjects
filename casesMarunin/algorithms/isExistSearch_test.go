package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsExistSort(t *testing.T) {
	expected := false
	expected2 := true
	actual := IsExistSort([]int{1000, 679, 100, 52, 10, 0, -8, -56, -100, -1000, -2000}, 1001)
	actual2 := IsExistSort([]int{-2000, -1000, -100, -56, -8, 0, 10, 52, 100, 679, 1000}, 679)
	assert.Equal(t, expected, actual)
	assert.Equal(t, expected2, actual2)
}
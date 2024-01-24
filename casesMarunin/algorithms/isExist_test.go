package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsExist(t *testing.T) {
	expected := false
	expected2 := true
	actual := IsExist([]int{1000, 679, 100, 52, 10, 0, -8, -56, -100, -1000, -2000}, 98)
	actual2 := IsExist([]int{1000, 679, 100, 52, 10, 0, -8, -56, -100, -1000, -2000}, 1000)
	assert.Equal(t, expected, actual)
	assert.Equal(t, expected2, actual2)
}
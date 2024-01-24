package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortStrings(t *testing.T) {
	expected := []string{"Иван Браташ", "Андрей Ворона", "Александр Кузнецов", "Эмиль Кузнецов", "Дмитрий Фефелов", "Обид Ясинов"}
	actual := SortStrings([]string{"Дмитрий Фефелов", "Андрей Ворона", "Иван Браташ", "Обид Ясинов", "Эмиль Кузнецов", "Александр Кузнецов"})
	assert.Equal(t, expected, actual)
}

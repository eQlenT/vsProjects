package algorithms

import (
	"fmt"
	"sort"
)
//Напишите фукнцию которая принимает на вход слайс []int отсортированный по убыванию и переменную x. Возвращает true если x есть в слайсе. С использованием sort.Search
func IsExistSort(sortedSlice []int, x int) bool {
	i := sort.Search(len(sortedSlice), func(i int) bool { return sortedSlice[i] >= x })
	if i < len(sortedSlice) && sortedSlice[i] == x {
         return true
		} else {
		fmt.Printf("x=%d больше максимального элемента слайса\n", x)
		return false
	}
}
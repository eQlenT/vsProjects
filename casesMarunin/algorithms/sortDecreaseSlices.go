package algorithms

import (
	"slices"
	"sort"
)

// Напишите фукнцию которая принимает на вход слайс []int и возвращает его отсортированным по убыванию с использованием slices.Sort
func SortDecreaseSlices(s []int) []int {
	res := make([]int, len(s))
	slices.Sort(s)
	for i := range s {
		res[i] = s[len(s)-1-i]
	}
	return res
}

// //Напишите фукнцию которая принимает на вход слайс []int и возвращает его отсортированным по убыванию с использованием sort.Slice
func SortDecreaseSort(s []int) []int {
	sort.Slice(s, func(i, j int) bool { return s[i] > s[j] })
	return s
}

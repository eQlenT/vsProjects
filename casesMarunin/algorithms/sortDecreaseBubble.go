package algorithms
//Напишите фукнцию которая принимает на вход слайс []int и возвращает его отсортированным по убыванию 
//без использования sort.Slice и slices.Sort
func SortDecreaseBubble(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < (len(s) - 1 - i); j++ {
			if s[j] < s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}

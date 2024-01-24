package algorithms
//Напишите фукнцию которая принимает на вход слайс []int отсортированный по убыванию и переменную x. Возвращает true если x есть в слайсе. Без использования библиотечных функций
func IsExist(sortedSlice []int, x int) bool {
	searchMap := make(map[int]int, len(sortedSlice))
	for i, v := range sortedSlice {
		searchMap[v] = i
	}
	if _, ok := searchMap[x]; ok {
		return ok
	} else {
		return ok
	}
}
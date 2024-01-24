package main

import (
	algs "casesmarunin/algorithms"
	"fmt"
	"slices"
)

func main() {
	someSlice := []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
	sortedIncreaseSlice := someSlice // для функции isExistSort
	slices.Sort(sortedIncreaseSlice) // для функции isExistSort
	var slcToSort = []string{"Анна Иванова", "Оксана Федорова", "Наталья Потапова", "Владимир Кузьмин", "Роман Алексеев", "Виктория Алексеева", "Ирина Андреева", "Александр Смирнов", "Татьяна Волкова", "Борис Петров", "Лариса Захарова", "Мария Белова", "Николай Колесников", "Евгения Орлова", "Галина Морозова", "Семен Морозов", "Даниил Григорьев", "Кирилл Новиков", "Екатерина Сидорова", "Иван Федоров", "Павел Карпов", "Раиса Яковлева", "Светлана Павлова", "Михаил Орлов", "Юлия Смирнова", "Ульяна Иванова", "Яна Алексеева", "Юрий Смирнов", "Леонид Михайлов", "Ярослав Кузьмин"}
	fmt.Println(algs.SortDecreaseBubble(someSlice))
	fmt.Println(algs.SortDecreaseSlices(someSlice))
	fmt.Println(algs.SortDecreaseSort(someSlice))
	fmt.Println(algs.IsExist(algs.SortDecreaseSlices(someSlice), 8))
	fmt.Println(algs.IsExistSort(sortedIncreaseSlice, 8))
	fmt.Println(slcToSort)
	fmt.Println(algs.SortStrings(slcToSort))
	fmt.Println(algs.Fridays())
}

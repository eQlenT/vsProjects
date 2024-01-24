package main

import (
	"fmt"
	"sort"
	"strings"
)

var slcToSort = []string{"Дмитрий Фефелов", "Андрей Ворона", "Иван Браташ", "Обид Ясинов", "Эмиль Кузнецов", "Александр Кузнецов"}

// Напишите функцию которая сортирует слайс строк вида
// "<Имя> <Фамилия>" по фамилии, а потом по имени.
// Но на выходе должен быть слайс строк вида "<Имя> <Фамилия>"
type human struct {
	Name    string
	Surname string
}

// конструктор
func makeHuman(name string, surname string) human {
	return human{
		Name:    name,
		Surname: surname,
	}
}

func parseHuman(input string) human {
	tmp := strings.Split(input, " ")
	return makeHuman(tmp[0], tmp[1])
}

type humans []human

func (h humans) sort() {
	sort.Slice(h, func(i, j int) bool { return h[i].Name < h[j].Name })
	sort.Slice(h, func(i, j int) bool { return h[i].Surname < h[j].Surname })
}

func sortStrings(toSort []string) []string {
	var slc []human
	for _, pair := range toSort {
		slc = append(slc, parseHuman(pair))
	}
	humans(slc).sort()
	for i := range toSort {
		toSort[i] = slc[i].Name + " " + slc[i].Surname
	}
	return toSort
}

func main() {
	fmt.Println(sortStrings(slcToSort))
}

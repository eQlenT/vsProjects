package algorithms

import (
	"sort"
	"strings"
)

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

func SortStrings(slcToSort []string) []string {
	slc := make([]human, 0)
	for _, pair := range slcToSort {
		slc = append(slc, parseHuman(pair))
	}
	humans(slc).sort()
	for i, v := range slc {
		slcToSort[i] = v.Name + " " + v.Surname
	}
	return slcToSort
}

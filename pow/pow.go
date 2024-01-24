package main

import "fmt"

func pow(base, exp int) int {
	switch exp {
	case 0:
		return 1
	case 1:
		return base
	default:
		return base * (pow(base, exp-1))
	}
}
func main() {
	res := pow(2, 4)
	fmt.Println(res)
}

/*package main

import "fmt"

func main() {
    // сложение чисел:
    oneHundred := 100
    fiveHundred := 500
    fmt.Printf("%d + %d = %d\n", oneHundred, fiveHundred, oneHundred + fiveHundred)

    // сложение строк:
    oneHundredStr := "сто"
    fiveHundredStr := "пятьсот"
    fmt.Printf("%s + %s = %s\n", oneHundredStr, fiveHundredStr, oneHundredStr+fiveHundredStr)
}*/
/*package main

import "fmt"

func main() {
    russianAlphabet := []string{}

    var ch rune

    for ch = 'а'; ch <= 'я'; ch++ {
        russianAlphabet = append(russianAlphabet, string(ch))
    }

    fmt.Printf("%s — последняя буква в алфавите.\n", russianAlphabet[len(russianAlphabet)-1])
}


/* package main

import (
    "fmt"
    "sort"
    "strings"
    "time"
)

var database = map[string]string{
    "Сергей":  "Омск",
    "Соня":    "Москва",
    "Алексей": "Калининград",
    "Миша":    "Москва",
    "Дима":    "Челябинск",
    "Алина":   "Красноярск",
    "Егор":    "Пермь",
    "Коля":    "Красноярск",
    "Артём":   "Владивосток",
    "Петя":    "Михайловка",
}

var offsetUTC = map[string]int{
    "Москва":          3,
    "Санкт-Петербург": 3,
    "Новосибирск":     7,
    "Екатеринбург":    5,
    "Нижний Новгород": 3,
    "Казань":          3,
    "Челябинск":       5,
    "Омск":            6,
    "Самара":          4,
    "Ростов-на-Дону":  3,
    "Уфа":             5,
    "Красноярск":      7,
    "Воронеж":         3,
    "Пермь":           5,
    "Волгоград":       3,
    "Краснодар":       3,
    "Калининград":     2,
    "Владивосток":     10,
}

func whatTime(city string) string {
    utcTime := time.Now().UTC()
    friendTime := utcTime.Add(time.Duration(offsetUTC[city]) * time.Hour)
    return friendTime.Format("15:04")
}

func formatCountFriends(count int) string {
    if count == 1 {
        return "1 друг"
    }
    if count >= 2 && count <= 4 {
        return fmt.Sprintf("%d друга", count)
    }
    return fmt.Sprintf("%d друзей", count)
}

func processAlice(query string) string {
    if query == "сколько у меня друзей?" {
        return fmt.Sprintf("У тебя %s", formatCountFriends(len(database)))
    }
    if query == "кто все мои друзья?" {
        var friends []string
        for name := range database {
            friends = append(friends, name)
        }
      // сортируем друзей по алфавиту
        sort.Strings(friends)
        return fmt.Sprintf("Твои друзья: %s", strings.Join(friends, ", "))
    }

    if query == "где все мои друзья?" {
        uniqueCities := make(map[string]int)
        // заполняем мапу без дублирования городов
        for _, city := range database {
            uniqueCities[city] = 1
        }
        var cities []string
        // получаем уникальные названия городов
        for city := range uniqueCities {
            cities = append(cities, city)
        }

     // сортируем города по алфавиту
        sort.Strings(cities)
        return fmt.Sprintf("Твои друзья в городах: %s", strings.Join(cities, ", "))
    }
    return "неизвестный запрос"
}

func processFriend(name, query string) string {
    city, found := database[name]
    if found {
        if query == "ты где?" {
            return fmt.Sprintf("%s в городе %s", name, city)
        }
		if query == "который час?"{// добавьте проверку запроса - который час?
            city, found := offsetUTC[city]
            fmt.Println(city)
			if !found { // проверьте, существует ли город в offsetUTC
				return fmt.Sprintf("Не могу определить время в городе %s", city) // если нет, то верните строку:
																					// "Не могу определить время в городе <название_города>"
			} else {
				return fmt.Sprintf("Там сейчас %s", whatTime(city)) // если город есть, то вызовете whatTime() и
																	// верните "Там сейчас <время>"
			}
        }
        return "неизвестный запрос"
    }
    return fmt.Sprintf("У тебя нет друга по имени %s", name)
}

func processQuery(query string) string {
    elements := strings.Split(query, ", ")
    if len(elements) != 2 {
        return "неизвестный запрос"
    }
    if elements[0] == "Алиса" {
        return processAlice(elements[1])
    }
    return processFriend(elements[0], elements[1])
}

func runner() {
    queries := []string{
        "Алиса, сколько у меня друзей?",
        "Алиса, кто все мои друзья?",
        "Алиса, где все мои друзья?",
        "Алиса, кто виноват?",
        "Коля, ты где?",
        "Соня, что делать?",
        "Антон, ты где?",
        "Алексей, который час?",
        "Артём, который час?",
        "Антон, который час?",
        "Петя, который час?",
    }
    for _, query := range queries {
        fmt.Println(query, "-", processQuery(query))
    }
}

func main() {
    runner()
}
*/
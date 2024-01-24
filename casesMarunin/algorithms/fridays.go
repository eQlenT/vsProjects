package algorithms

import (
	"fmt"
	"time"
)

func Fridays() string {
	weekdayCount := map[time.Weekday]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
	}
	date := time.Date(1582, 10, 15, 0, 0, 0, 0, time.UTC) //начало последнего завершившегося 400-летнего цикла
	for ; date != time.Date(1982, 10, 15, 0, 0, 0, 0, time.UTC); date = date.AddDate(0, 0, 1) {
		if date.Day() == 13 {
			weekdayCount[date.Weekday()]++
		}
	}
	counter := 0 
	for k, v := range weekdayCount {
		if v>weekdayCount[time.Friday]{
			counter = weekdayCount[k]
		}
	}
	if counter == 0 {
		fmt.Println(weekdayCount)
		return "ЧТД"
	} else {
	fmt.Println(weekdayCount)
	return "Не доказано"
	}
}

package main

import (
	"fmt"
	"strings"
)


func main() {
	var a int
	table(jsonWork("info.json"))
	fmt.Scan(&a)
}

func maxLenghtInColumn(dict map[int][]string) map[int]int {
	var mapResult = make(map[int]int)

	for i := 0; i < len(dict); i++ {
		for j := 0; j < len(dict[i]); j++ {
			if mapResult[j] < len([]rune(dict[i][j])) {
				mapResult[j] = len([]rune(dict[i][j]))
			}
		}
	}

	return mapResult
}

func table(userData map[int][]string) {
	maxLenght := maxLenghtInColumn(userData)
	row(userData[0], true, maxLenght)
	for i := 1; i < len(userData); i++ {
		row(userData[i], false, maxLenght)
	}

}

func row(name []string, isTitle bool, maxLenghtInColumn map[int]int) {
	var strTop string
	var strMiddle string
	var strBottom string
	var topTitle string
	for i := 0; i < len(name); i++ {

		topTitle += fmt.Sprintf("%s", strings.Repeat("_", maxLenghtInColumn[i]+5))
		strTop += fmt.Sprintf("|  %s  ", strings.Repeat(" ", maxLenghtInColumn[i]))

		strMiddle += fmt.Sprintf("|  %s%s  ", name[i], strings.Repeat(" ", maxLenghtInColumn[i]-len([]rune(name[i]))))

		strBottom += fmt.Sprintf("|__%s__", strings.Repeat("_", maxLenghtInColumn[i]))

	}
	if isTitle {
		fmt.Printf("%s_\n", topTitle)
	}
	fmt.Printf("%s|\n", strTop)
	fmt.Printf("%s|\n", strMiddle)
	fmt.Printf("%s|\n", strBottom)

}

// 0: {"Id", "Name", "Age", "IsAdmin", "Location", "remember"},
// 1: {"1", "Egor", "Age", "true", "1", "232"},
/*
	// КОД ДЛЯ ВВОДА ДАННЫХ ЧЕРЕЗ КОНСОЛЬ, НО ЕЩЕ ДОРАБОТАТЬ НАДО ЧУТКА
	var data []string
	var dataTitle []string
	var temp string

	fmt.Println("Введите наименования полей через пробел (для продолжения напишите 'exit') ")
	for {
		fmt.Scan(&temp)
		if temp == "exit" {
			userData[0] = dataTitle
			break
		}
		dataTitle = append(dataTitle, temp)
	}
	fmt.Println("Введите данные полей через пробел (для выхода напишите 'exit') ")
	for {
		fmt.Scan(&temp)
		if temp == "exit" {
			userData[1] = data
			break
		}
		data = append(data, temp)
		if len(data) == len(dataTitle) {
			userData[1] = data
			break
		}

	}

*/


// "title": ["id", "name", "age", "isAdmin", "location"],
//     "data": {
//         "1": {
//             "id": 1,
//             "name": "egor",
//             "age": 23,
//             "isAdmin": true,
//             "location": "samara"
//         },
//         "2": {
//             "id": 2,
//             "name": "sas",
//             "age": 19,
//             "isAdmin": false,
//             "location": "Moscow"
//         }
//     }

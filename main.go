/*Изучение массивов, цикла for и срезов*/

package main

import (
	"fmt"
)

func main() {
	var todoList = [3]string {
		"Изучаем массивы",
		"Изучить срезы",
		"Изучить цикл for",
	}

	for index, item := range todoList {
		fmt.Printf("%d. %s\n", index, item)
	}
}

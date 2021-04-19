package main

import (
	"fmt"
)

func main() {
	var todoList = [4]string {
		"Изучаем массивы",
		"Изучить срезы",
		"Изучить цикл for",
		"Добавить изменения в Git",
	}

	for index, item := range todoList {
		fmt.Printf("%d. %s\n", index, item)
	}
}

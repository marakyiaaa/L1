package main

import (
	"fmt"
)

func quickSort2(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	var (
		less []int
		more []int
	)

	pivot := arr[len(arr)-1] // Выбираем последний элемент как опорный

	for _, value := range arr[:len(arr)-1] {
		if value <= pivot {
			less = append(less, value)
		} else {
			more = append(more, value)
		}
	}

	// Рекурсивно сортируем меньшие и большие элементы
	less = quickSort2(less)
	more = quickSort2(more)

	// Объединяем отсортированные меньшие элементы, опорный элемент и отсортированные большие элементы
	return append(append(less, pivot), more...)
}

func main() {
	arr := []int{54, 4, -5, 2, 79, 4}
	fmt.Println(quickSort2(arr))
}

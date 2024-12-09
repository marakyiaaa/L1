package main

import (
	"fmt"
)

/*через значение опрного*/

func quickSort2(arr []int) []int {
	if len(arr) <= 1 {
		return arr //завершаем рекурсию при 1 оставшемся элементе
	}

	var (
		less []int
		more []int
	)

	elem := arr[len(arr)-1] // Выбираем (значение) последний элемент как опорный

	for _, value := range arr[:len(arr)-1] { //проходимся по подсрезу,невкл. в себя последний элемент
		if value <= elem {
			less = append(less, value)
		} else {
			more = append(more, value)
		}
	}

	// Рекурсивно сортируем меньшие и большие элементы
	less = quickSort2(less)
	more = quickSort2(more)

	return append(append(less, elem), more...)
}

func main() {
	arr := []int{54, 4, -5, 2, 79, 4}
	fmt.Println(quickSort2(arr))
}

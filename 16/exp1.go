package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr // когда массив имеет 0 или 1 элементов он уже отсортирован => прекращаем рекурсию
	}
	var (
		less []int
		more []int
	)
	elem := len(arr) - 1           //берем опорным элементов последний элемент массива
	more = append(more, arr[elem]) //запись опорного элемента в начала моссива бОльших чисел

	for c := 0; c < elem; c++ {
		if arr[c] <= arr[elem] {
			less = append(less, arr[c]) //запись меньших значеий в массив с меньшими числами
		} else {
			more = append(more, arr[c]) //запись больших значиний в массив с большими числами
		}
	}
	return append(quickSort(less), quickSort(more)...)
}

func main() {
	arr := []int{54, 4, -5, 2, 79, 4}
	fmt.Println(quickSort(arr))
}

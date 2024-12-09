package main

import "fmt"

/*через индекс опрного и запись его в больщий массив*/

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr // когда массив имеет 0 или 1 элементов он уже отсортирован => прекращаем рекурсию
	}
	var (
		less []int
		more []int
	)
	elem := len(arr) - 1           //берем индкс опорного элемента - последний элемент массива
	more = append(more, arr[elem]) //запись опорного элемента в начала моссива бОльших чисел

	for c := 0; c < elem; c++ {
		if arr[c] <= arr[elem] {
			less = append(less, arr[c]) //запись меньших значеий в массив с меньшими числами
		} else {
			more = append(more, arr[c]) //запись больших значиний в массив с большими числами
		}
	}
	return append(quickSort(less), quickSort(more)...)
	//... - здесь используется для распаковки элементов среза, чтобы передать их как отдельные аргументы в функцию append.
}

func main() {
	arr := []int{54, 4, -5, 2, 79, 4}
	fmt.Println(quickSort(arr))
}

package main

import "fmt"

/* Если элемент найден, функция возвращает его индекс; если не найден, возвращается -1*/

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1 //границы массива

	for left <= right {
		middle := (left + right) / 2 //вычисляем средний индекс
		if target > arr[middle] {
			left = middle + 1
		} else if target < arr[middle] {
			right = middle - 1
		} else {
			return middle
		}
	}

	return -1
}
func main() {
	arr := []int{-5, 4, 12, 26, 50, 76}
	fmt.Println(binarySearch(arr, 50))
	fmt.Println(binarySearch(arr, 3))
}

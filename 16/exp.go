package main

import "fmt"

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort1(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort1(arr, low, p-1)
		arr = quickSort1(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort1(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{2, 17, 27, 11, -9, 14}
	fmt.Println(quickSortStart(arr))
}

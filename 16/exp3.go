package main

import (
	"fmt"
)

/*перезапись базового массива*/

func quickSort3(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSort3(arr, low, pivotIndex-1)
		quickSort3(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func main() {
	arr := []int{54, 4, -5, 2, 79, 4}
	fmt.Println(arr)
	quickSort3(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

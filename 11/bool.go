package main

import "fmt"

func intersectionsBool(set1, set2 []int) []int {
	elemSetOne := make(map[int]bool) //для хранения элементов set1

	for _, elem := range set1 {
		elemSetOne[elem] = true
	}

	result := []int{}

	for _, elem := range set2 {
		if _, ok := elemSetOne[elem]; ok { //если elem присутвует так же в elemSetOne
			result = append(result, elem) //то записываем его в срез
			delete(elemSetOne, elem)      //и сразу удаляем,чтобы не было дубликации
		}
	}
	return result
}

func main() {
	set1 := []int{66, 52, 3, 4}
	set2 := []int{3, 52, 5, 6}
	fmt.Println("Пересечение множеств:", intersectionsBool(set1, set2))
}

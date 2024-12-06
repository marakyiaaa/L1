package main

import "fmt"

/*
Множество A: {1, 2, 3, 4}
Множество B: {3, 4, 5, 6}
Пересечение: {3, 4}
*/

func intersectionsStruct(set1, set2 []int) []int {
	//добавляем проверку длины для оптимизации
	var smallerSet, largerSet []int
	if len(set1) < len(set2) {
		smallerSet, largerSet = set1, set2
	} else {
		smallerSet, largerSet = set2, set1
	}

	elemSetOne := make(map[int]struct{}) //для хранения элементов set1

	for _, elem := range smallerSet {
		elemSetOne[elem] = struct{}{}
	}

	result := []int{}

	for _, elem := range largerSet {
		if _, ok := elemSetOne[elem]; ok { //если elem присутвует так же в elemSetOne
			result = append(result, elem) //то записываем его в срез
			delete(elemSetOne, elem)      //и сразу удаляем,чтобы не было дубликации
		}
	}
	return result
}

func main() {
	set1 := []int{1, 2, 3, 4}
	set2 := []int{3, 4, 5, 6}
	fmt.Println("Пересечение множеств:", intersectionsStruct(set1, set2))
}

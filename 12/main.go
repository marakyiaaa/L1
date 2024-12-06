package main

import "fmt"

func mySet(consistency []string) map[string]struct{} {
	intersection := make(map[string]struct{})

	//ключем является слово - struct{}{} чтобы не занимать память
	for _, elem := range consistency {
		intersection[elem] = struct{}{}
	}
	return intersection
}

func main() {
	consistency := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("Собственное множество:", mySet(consistency))
}

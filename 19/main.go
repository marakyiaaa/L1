package main

import "fmt"

func reverse(str string) string {
	runeStr := []rune(str)

	for i, j := 0, len(runeStr)-1; i < j; i, j = i+1, j-1 {
		runeStr[i], runeStr[j] = runeStr[j], runeStr[i]
	}
	return string(runeStr)
}
func main() {
	str := "главрыба"
	fmt.Println(reverse(str))
}

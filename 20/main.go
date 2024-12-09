package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s) // Разбиваем строку на слова

	// Переворачиваем массив слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ") // Объединяем слова обратно в строку
}

func main() {
	input := "snow dog sun"
	fmt.Println(reverseWords(input))
}

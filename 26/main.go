package main

import (
	"fmt"
	"strings"
)

func unique(str string) bool {
	m := make(map[int]struct{})

	for _, val := range strings.ToLower(str) {
		m[int(val)] = struct{}{}
	}
	return len(m) == len(str)

}
func main() {
	str := []string{"abcd", "abCdefAaf", "aabcd"}

	for i := range str {
		fmt.Println(unique(str[i]))
	}
}

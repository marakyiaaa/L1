package main

import "fmt"

func main() {
	//можем присвоить ему любой тип пустому интерфейсу
	var emptyInterface interface{} = true

	switch x := emptyInterface.(type) {
	case string:
		fmt.Println("string:", x)
	case int:
		fmt.Println("int:", x)
	case bool:
		fmt.Println("bool:", x)
	case chan int:
		fmt.Println("chan int", x)
	case chan string:
		fmt.Println("chan string", x)
	case chan bool:
		fmt.Println("chan bool", x)
	default:
		fmt.Print("I don't know")
	}
}

package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Action struct {
	Human //встраивание родительской структуры => все методы родителя видны Action
}

func (h *Human) Hello() {
	fmt.Printf("Hello, I`m %s\n", h.name)
}

func (h *Human) InfoHuman() {
	fmt.Printf("I'm %d years old\n", h.age)
}

//переопределение метода в Action блокирует методы родителя
//func (a *Action) Hello() {
//	fmt.Println("Hello, I`m Den")
//}

func main() {
	human := Action{
		Human{name: "Anton",
			age: 5},
	}
	human.Hello()
	human.InfoHuman()
}

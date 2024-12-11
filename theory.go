package main

/*//1) Какой самый эффективный способ конкатенации строк?

// самый базовый способ - Оценка производительности: 67.8%.
func concat2operator(str1, str2 string) string {
	return str1 + str2
}

// strings.Builder - Оценка производительности: 80.5% (+12.7).
// Это аналог bytes.Buffer, но при вызове метода String() не происходит повторного выделения памяти и копирования данных.
func concat2builder(str1, str2 string) string {
	var builder strings.Builder
	builder.Grow(len(str1) + len(str2)) //выделение памяти под длину соедененных строк,чтобы не было нескольких реаллокаций буфера.
	builder.WriteString(str1)
	builder.WriteString(str2)
	return builder.String()
}

func main() {
	str1 := "lalala"
	str2 := "banbanban"
	fmt.Println(concat2operator(str1, str2))
	fmt.Println(concat2builder(str1, str2))
}*/

/*
// 7) В какой последовательности будут выведены элементы map[int]int?
// Пример:
// m[0]=1
// m[1]=124
// m[2]=281

	func main() {
		m := map[int]int{
			0: 1,
			1: 124,
			2: 281,
		}

		for key, value := range m {
			fmt.Println(key, value)
		}
	}
*/

/* //10)
func update(p *int) {
	b := 2
	p = &b
	fmt.Println(*p) //2
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p) //1
	update(p)
	fmt.Println(*p) //1
}*/

/*//11)

func main() {
	wg := sync.WaitGroup{} //создаем WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1) //запускаем счетчик
		go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done() //счетчик --
		}(&wg, i)
	}

	wg.Wait()
	fmt.Println("exit")
}*/

/*// 12)
func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n) //0
}
*/

/*// 13)
func someAction(v []int8, b int8) {
	v[0] = 100       //{100, 2, 3, 4, 5}
	v = append(v, b) //{100, 2, 3, 4, 5, 6}
	fmt.Println(v)
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}
*/

/*//14)

func main() {
	slice := []string{"a", "a"} //len 2, cap 2 -> m1

	func(slice []string) { //копия
		slice = append(slice, "a") //{"a", "a", "a"} len 3, cap 4 -> m2
		slice[0] = "b"             //{"b", "a", "a"} len 3, cap 4 -> m2
		slice[1] = "b"             //{"b", "b", "a"} len 3, cap 4 -> m2
		fmt.Print(slice)           //{"b", "b", "a"} len 3, cap 4 -> m2
	}(slice)

	fmt.Print(slice) //{"a", "a"}
}*/

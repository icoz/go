package main

import "fmt"

const size uint = 3

func main() {
	// инициализируется значениями по-умолчанию; для int - 0
	// при создании сразу алоцируется память
	var a1 [3]int
	fmt.Println("Array", a1, "length", len(a1))

	// можно использовать типизированную беззнаковую константу
	var a2 [2 * size]bool
	fmt.Println(a2, "length", len(a2))

	// ... используется длина при инициализации
	a3 := [...]int{1, 2, 3}
	fmt.Println("length at initialization", a3, "length", len(a3))

	a3[1] = 12
	fmt.Println("after change", a3)

	// a3[4] = 12 нельзя, проверка при компиляции
	// invalid array index 4 (out of bounds for 3-element array)

	var aa [3][3]int
	aa[1][1] = 1
	fmt.Println("array of arrays", aa)

	/*
		массивы это очень низкоуровневый тип данных; длина массива
		является частью его типа; это значить что в функцию которая ждет
		массив их трех элементов нельзя передать массив другой длины, пото-
		му что длина массива - это его тип.
	*/

}

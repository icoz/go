package main

// стандариный пакет для форматирования ввода
import "fmt"

/*
	slice
		len 1
		capacity 2
		array [100]

	slice это структура данных к которой есть длина, capacity и ссылка
	на массив с данными; если массив заполнен, будет создаваться другой массив
	и увеличиваться capacity c копированием всех значений в новый массив
*/

func main() {
	var sl []int
	fmt.Println("empty slice", sl)
	sl = append(sl, 100)
	fmt.Println("not empty", sl)
	fmt.Println("slice's length", len(sl))

	fmt.Println("internal array's length in slice", sl, cap(sl))

	// adding to slice always x2 to capacity
	sl = append(sl, 102)
	fmt.Println("internal array's length in slice", sl, cap(sl))
	sl = append(sl, 103)
	sl = append(sl, 104)
	fmt.Println("internal array's length in slice", sl, cap(sl))
	sl = append(sl, 105)
	fmt.Println("internal array's length in slice", sl, cap(sl))

	// short initialization
	sl2 := []int{10, 20}
	sl2 = append(sl2, 30)
	fmt.Println(sl2, len(sl2), cap(sl2))

	// добавить слайс в слайс
	// нужен оператор ... для приведения типа
	// тк ожидает int
	sl = append(sl, sl2...)
	fmt.Println(sl)

	// можно так
	var slm [][]int
	slm = append(slm, sl2)
	fmt.Println(slm)

	// создать слайс с нужной длиной сразу
	// и иницилизацией
	slice3 := make([]int, 10)
	slice3 = append(slice3, 1)
	fmt.Println(slice3, len(slice3), cap(slice3))

	// создать слайс с нужной длиной и размером
	// slice4 := make([]int, 10, 15)
	slice4 := make([]int, 10, 15)
	fmt.Println(slice4, len(slice4), cap(slice4))
	slice4 = append(slice4, []int{1, 2, 3, 4, 5, 6}...)
	fmt.Println(slice4, len(slice4), cap(slice4))

	// внутри слайса - ссылка на массив; она копируется если просто присвоить
	// ссылка на одну и ту же область данных
	slice5 := slice4
	slice5[1] = 100500
	fmt.Println(slice4, slice5)

	// but, if...
	slice4 = append(slice4, []int{1, 2, 3, 4, 5, 6}...)
	slice4[1] = 999
	fmt.Println(slice4, slice5)
	// here we append a bunch of values and slice create another array
	// which referenced to another piece of memory

	// неправильно - он попоробует скопировать сколько влезет
	var slice6 []int
	copy(slice6, slice5)
	fmt.Println(slice6) // []

	// правильное копирование слайса - создание нужной capacity --> нужным размером
	// те сколько нужно для копирования
	// копирование не append, оно заменяет все значения
	slice7 := make([]int, len(slice5), len(slice5))
	// slice7 := make([]int, 2, 2) // скопирует только первые два
	copy(slice7, slice5)
	fmt.Println(slice7)

	fmt.Println("часть слайса", slice7[1:5], slice7[:2], slice7[10:])
	slice8 := append(slice7[:2], slice7[10:]...)
	fmt.Println("из кусков слайса", slice8)

	// создать слайс из массива
	// здесь ссылаестя на ту же обдасть в памяти
	a := [...]int{5, 6, 7}
	sl8 := a[:]
	a[1] = 8
	fmt.Println("slice from array", sl8)

	sl9 := make([]int, 10)
	copy(sl9[3:6], []int{1, 2, 4})
	// println(sl9)
	sl9 = append(sl9[3:6], []int{1, 2, 4}...)
	// println печатает length/capacity и ссылку в памяти на массив
	// форматированный ввод/вывод
	fmt.Printf("slice elements is %v\nnewline\n", sl9)
}

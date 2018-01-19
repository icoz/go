package main

import "fmt"

/*
	map - структура данных - она же hash - она же ассоциативный массив
	если, например, к вас тип int и индекс 1000, то вам не нужно держать
	как в массиве 1000 элементов, а только одно.
*/

func main() {
	// тип ключа и тип данных
	var mm map[string]string
	// hash третьего уровня
	// var mm2 map[string]map[string]map[string]string

	// panic: assigment to entry in nil map
	// nil - специальный тип данных в го, который говорит
	// что там ничего нет, те ни на что не указывает
	// mm["test"] = "false"

	fmt.Println("uninitialized map", mm)

	// нужна полная инициализация
	// var mm2 map[string]string = map[string]string{}
	// or short
	mm2 := map[string]string{}
	mm2["test"] = "ok"
	fmt.Println(mm2)

	// или через make которая инициализирует map пустой
	var mm3 = make(map[string]string)
	mm3["firstName"] = "Harry"
	fmt.Println("mm3", mm3)

	// получение значения (короткое синтаксиc объявление)
	firstName := mm3["firstName"]
	fmt.Println("fisrtName", firstName, len(firstName))

	// если обратиться к несуществующему ключу - вернется значение по-умолчанию
	lastName := mm3["lastName"]
	fmt.Println("lastName", lastName, len(lastName))

	// проверка на то, что значение есть
	// вторым значением map возвращает trur/false
	lastName, ok := mm3["lastName"]
	fmt.Println("lastName is", lastName, "exist", ok)

	// получение только признака существования
	_, exist := mm3["firstName"]
	fmt.Println("fistName existence:", exist)

	// удаление значения
	key := "firstName"
	// delete(mm3, "firstName")
	delete(mm3, key)
	_, exist = mm3["firstName"]
	fmt.Println("fistName existence:", exist)

	// map это значение которое присваивается по ссылке
	// вы не копируете map, там глубже лежит какая-то структура данных
	// стандартного функционала для копирования map нет
	// нужно копировать в цикле или рекурсивно
	mm4 := mm3
	mm4[key] = "test"
	fmt.Println(mm3, mm4)

	// порядок ключей в map недетерминирован; всегда случайный порядок
	// в php в массиве/map порядок фиксирован
}

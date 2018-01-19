package main

const someInt = 1        // нетипизировання константа
const typeInt int32 = 17 // типизированная
const fullName = "Harry"

// короткий синтаксис
const (
	flagKey1 = 1
	flagKey2 = 2
)

const (
	zero = iota
	one  = iota + 1 // автоинкремент, можно проводить операции
	two
	_     // пустая переменная, пропуск iota
	floor // = 4
)

const (
	_         = iota // пропускаем первое значение
	KB uint64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	// ошибка переполнение типа
	// ZB
)

func main() {
	println(zero, one, two, floor)
	// println(KB, MB, GB, TB, PB, EB)
}

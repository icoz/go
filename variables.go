package main

// import 'fmt'

func main() {

	// целые числа
	var i int = 10 // платформозависимый тип, 32 или 64 бита
	var autoInt = -10
	var bigInt int64 = 1<<32 - 1          // int8, int16, int32, int64
	var unsignedInt uint = 100500         // платформозависимый тип, 32 или 64 бита
	var unsignedBigInt uint64 = 1<<64 - 1 // uint8, uint16, uint32, uint64

	println("integers", i, autoInt, bigInt, unsignedInt, unsignedBigInt)

	// числа с плавающей точкой (одинарной, двойной точности)
	var p float32 = 3.14 // float = float32, float64
	println("float: ", p)

	// булевые переменные
	var b = true
	println("bool variable: ", b)

	// строки (встроены в сам язык и они все в уникоде)
	// удобно работать
	var hello string = "Hello\n\t"
	var world = "World"
	println(hello, world)

	// бинарные данные
	// байт - алиас для uint8 - он содержит 1 байт в памяти - 8 бит
	// это какой-то набор данных для представление ... чего-то
	var rawBinary byte = '\x27' // по hex коду; одинарные кавычки; двойные для строк; одинарные для указания конкретного символа
	println("rawBinary: ", rawBinary)

	// короткое объявление; это блихко к динасическим языкам где не нужно объявление переменных
	meaningOfLive := 42
	println("Meaning of live is:", meaningOfLive)

	// приведеине типо
	println("float to int conversion:", int(p))
	var u1 uint = 17
	var s1 int = 23
	println(s1 + int(u1))
	println("int to string conversion", string(48)) // номер символа а не строчка

	/*
		комплексные числа // в го есть встроенная поддержка
	*/
	z := 2 + 3i
	println("complex number:", z)

	// escaping / апострофы
	// когда не нужно заботиться об экранировании символов
	escaping := `Hello\r\n
	World`
	println("as-is escaping:", escaping)

	/*
		значения по-умолчанию
		в го у всех значений есть значения по-умолчанию
	*/
	var defaultInt int
	var defaultFloat float32
	var defaultString string
	var defaultBool bool
	println("default values:", defaultInt, defaultFloat, defaultString, defaultBool)

	/*
		несколько переменных
		в го нет глобальных переменных
		есть переменные уровня пакета (package)
		в динамических языках вы платите за динамичность переменных памятью
		там есть структура переменной, внутри тип этой переменной и сразу
		залоцирована память для сразу всех возможных типов в ней

		в го и во всех компилируемых языках вы указавая переменные
		не получаете этог overhead, накладных расходов по памяти

		в php (5) например overhead 72 байта на переменную; в php (7) ок. 30
		в perl ок. 40 + расходы на все проверки интерпретатором относительно того,
		какой тип достать, как сложить и можно ли вообще и проч.
		и это все во время выподнения программы

		если у вас язык компилируемый, со статической типизацией, у вас все эти
		проверки выполняются один раз при компиляции; а во время работы программы
		этих проверок нет, если мы убедились что все корректно во вреся компиляции
		(указатели и ансейвы исключения)
	*/

	var v1, v2 string = "v1", "v2"
	println(v1, v2)
	var (
		m0 int = 12
		m1     = "string"
		m2     = 23
	)
	println(m0, m1, m2)
}

package main

import "fmt"

/*
	в го юникод встроен в сам язык; это не внешняя библиотека,
	которая оперирует над какими-то бинарными данными
	тип rune, которой указывается конкретное значение
*/

func main() {
	var symbol rune = 'a'
	var autoSymbol = 'a'
	// unicodeSymbol := 'put your unicodeSymbol here'
	unicodeSymbolByNumber := '\u2318'
	println(symbol, autoSymbol, unicodeSymbolByNumber)

	str1 := "Привет, Мир!"
	fmt.Println("ru:", str1, len(str1))
	for index, runeValue := range str1 {
		fmt.Printf("%#U at position %d\n", runeValue, index)
	}
	fmt.Println(str1[1])

	// юникод это многобайтная кодировка: символ занимает не один байт, а два
	// в один байт укладывается только одна ASCI таблица (управляющие символы,
	// немного ASCI графики и латинский алфавит); русских символов в ASCI таблице
	// нет
	// в китайском символ занимает три байта
	// если печатать через println, возвращаться будет уже байтовое значение ('П' = 159)

	// итерация по самим байтам
	// нужно преобразоваться в slice-byte
	// внутри строки этот слайс-байт и лежит
	// и в нем уже можно итерироваться как в слайсе
	bin := []byte(str1)
	fmt.Println("binary cn:", bin, len(bin))
	for idx, val := range bin {
		fmt.Printf("raw binary idx: %v, oct: %v, hex: %x\n", idx, val, val)
	}
}

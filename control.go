package main

func main() {
	println("управляющие структуры")

	a := true
	if a {
		println("hello, world")
	}
	b := 1
	c := 22
	if (b == 1 && a) || c != 22 {
		println("неявное преобразование (if b) не работает")
	}

	// в if может быть использована ;
	// оператор может быть сложным
	mm := map[string]string{"firstName": "Harry", "lastName": "Potter"}
	// // firstName и ok доступны только в области видимости if
	// if firstName, ok := mm["firstName"]; firstName != "Harry" {
	// 	println("firstName key exist =", firstName, ok)
	// } else {
	// 	println("no fisrtName")
	// }

	// if _, ok := mm["firstName"]; ok {
	// 	println("firstName key exist =", ok)
	// } else {
	// 	println("no fisrtName")
	// }

	if firstName, ok := mm["firstName"]; ok {
		println("firstName key exist =", firstName)
	} else {
		println("no fisrtName")
	}

	if firstName, ok := mm["firstName"]; !ok {
		println("no firstName")
	} else if firstName == "Harry" {
		println("firstName is Harry")
	} else {
		println("firstName is not Harry")
	}

	// в го только один цикл for
	for {
		println("бесконечный цикл while(true)")
		break
	}

	sl := []int{3, 4, 5, 6, 7, 8}
	value := 0
	idx := 0

	// операция по slice
	for idx < 4 {
		if idx < 2 {
			idx++
			continue
		}
		value = sl[idx]
		idx++
		println("while-c-type loop, idx:", idx, "value", value)
	}

	for i := 0; i < len(sl); i++ {
		println("c-style loop", i, sl[i])
	}

	for idx := range sl {
		println("range slice by index", idx)
	}

	for idx, val := range sl {
		println("range slice by index-val", idx, val)
	}

	for _, val := range sl {
		println("range slice by val and index-continue", val)
	}

	for key := range mm {
		println("range map by key", key)
	}

	for key, val := range mm {
		println("range map by key-val", key, val)
	}

	for _, val := range mm {
		println("range map by val and index-continue", val)
	}

	// switch
	mm["firstName"] = "Harry"
	mm["flag"] = "ok"
	switch mm["firstName"] {
	case "Ron", "Harry":
		println("switch - name is Harry")
		// в отличие от других языков не переходим в другой вариант по-умолчанию
	case "Petr":
		if mm["flag"] == "ok" {
			break // выходим из цикла, чтобы не выполнять переход в другой вариант
		}
		println("switch - name is Petr")
		fallthrough // переходим в следующий вариант
	default:
		println("switch - nobody")
	}

	// как замена множественных if else
	switch {
	case mm["firstName"] == "Harry":
		println("switch2 - Harry")
	case mm["lastName"] == "Dumbledore":
		println("switch2 - Dumbledore")
	}

	// если switch в цикле, для выхода из цикла используется метка/label
MyLoop:
	for key, val := range mm {
		println("switch in loop", key, val)
		switch {
		case key == "firstName" && val == "Harry":
			println("switch - we found Potter, break loop here")
			break MyLoop
		default:
			println("switch - nobody")
		}
	}
}

package main

import "fmt"
import "strings"
import "sort"

func main() {
	// input := map[int]string{
	// 	9:  "Сентябрь",
	// 	1:  "Январь",
	// 	2:  "Февраль",
	// 	10: "Октябрь",
	// 	5:  "Май",
	// 	7:  "Июль",
	// 	8:  "Август",
	// 	12: "Декарь",
	// 	3:  "Март",
	// 	6:  "Июнь",
	// 	4:  "Апрель",
	// 	11: "Ноябрь",
	// }
	// input2 :=
	// 	map[int]string{
	// 		10: "Зима",
	// 		30: "Лето",
	// 		20: "Весна",
	// 		40: "Осень",
	// 	}

	// GetMapValuesSortedByKey(input)
	// GetMapValuesSortedByKey(input2)

	// arr := []int{17, 23, 100500}
	// sl1 := []float32{1.1, 2.1, 3.1}
	// sl2 := []int32{4, 5}
	// str := IntSliceToString(arr)
	// MergeSlices(sl1, sl2)
	// msl := MergeSlices(sl, sl2)
	// fmt.Println(str)
	// fmt.Println(msl)
	// fmt.Println(int(sl[0]))
}

func ReturnInt() int {
	return 1
}

func ReturnFloat() float32 {
	return 1.1
}

func ReturnIntArray() [3]int {
	a := [...]int{1, 3, 4}
	return a
}

func ReturnIntSlice() []int {
	a := []int{1, 2, 3}
	sl := make([]int, len(a), len(a))
	copy(sl, a)
	return sl
}

func IntSliceToString(sl []int) string {
	str := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sl)), ""), "[]")
	return str
}

func MergeSlices(sl1 []float32, sl2 []int32) []int {
	mergeslice := make([]int, 5, 5)
	for idx, val := range sl1 {
		mergeslice[idx] = int(val)
	}
	for idx, val := range sl2 {
		mergeslice[idx+3] = int(val)
	}
	return mergeslice
}

func GetMapValuesSortedByKey(input map[int]string) []string {
	sortedarr := make([]string, len(input), len(input))
	var keys []int
	for key := range input {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for i, k := range keys {
		sortedarr[i] = input[k]
	}
	// listOfNums := strings.Join([]string{"3","2","1"}, ", ")
	// fmt.Println(sortedarr[0])
	return sortedarr
}

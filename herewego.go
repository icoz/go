package main

// import "math"
import "fmt"
import "time"
import "strconv"

// import "net/http"
// import "strconv"
// import "os"
// import "log"
// import "io/ioutil"
// import "strings"
// import "sort"
// import ("fmt"
// "strings"
// "sort"
// "os"
// "log"
// "io/ioutil"
// "strconv")
// godoc : package name / function
// godoc fmt Printliu

// variables inside a go is statically typed / type can't change once it's defined
// vars may be: ljljajlaksdll3242____

// func addThemUp(numbers []float64) float64{
// 	sum := 0.0
// 	for _, val := range numbers {
// 		sum += val
// 	}
// 	return sum
// }

// return two values from a function
// kind of neat; because most programming langs just
// don't allow you to do that
// func next2Values(number int) (int, int){
// 	return number+1, number+2
// }

// when undefined number of values passed
// func subtractThem(args ...int) int{
// 	finalValue := 0
// 	for _, value := range args {
// 		finalValue -= value
// 	}
// 	return finalValue
// }

// Факториа́л — функция, определённая на множестве неотрицательных целых чисел.
// Факториал натурального числа n определяется как произведение всех натуральных чисел от 1 до n включительно:
// you just to have to stare at it a little bit and eventually it just makes sense
// func factorial(num int) int{
// 	if num == 0 {
// 		return 1
// 	}
// 	return 	num * factorial(num - 1)
// }

// defer
// func printOne() { fmt.Println(1) }
// func printTwo() { fmt.Println(2) }

// func safeDiv(num1, num2 int) int{
// 	// error handler - recover
// 	// allow us to continue execution even
// 	// if a fatal error should occur
// 	defer func() {
// 		fmt.Println(recover())
// 		// recover()
// 	}()
// 	solution := num1 / num2
// 	return solution
// }

// func demPanic() {
// 	defer func() {
// 		fmt.Println(recover())
// 	}()

// 	panic("PANIC")
// }

// pointers
// here we pass value, not the reference to the value
// func changeXVal(x int) {
// 	x = 2
// }
// and here we pass reference of the value, not the value itself
// this aalow us to actually change the value at the memory address
// referenced by the pointer
// func changeXValNow(x *int) {
// 	// store 2 in the memmory address that x refers to
// 	// x has nothing to do with the value anymore
// 	// it actually deals with the actual variable
// 	*x = 2
// }

// func changeYValNow(yPtr *int) {
// 	*yPtr = 100
// }

// type Rectangle struct {
// 	leftX float64
// 	topY float64
// 	height float64
// 	width float64
// }

// method area for Rectangle struct
// func (rect *Rectangle) area() float64{
// 	return rect.width * rect.height
// }

// interfaces and usage a sort of polymorphism
// inside of go
// type Shape interface {
// 	area() float64
// }
// type Rectangle struct {
// 	height float64
// 	width float64
// }
// type Circle struct {
// 	radius float64
// }

// func (r Rectangle) area() float64 {
// 	return r.height * r.width
// }
// func (c Circle) area() float64 {
// 	return math.Pi * math.Pow(c.radius, 2)
// }
// func getArea(shape Shape) float64 {
// 	return shape.area()
// }

// ***** some bunch of string functions
// cover file i/o, accepting input from the user
// and casting when we'll create a web server
// also we cover go routines and channels

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World\n")
// }
// func handler2(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello Earth\n")
// }

// go routines
// func count(id int) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(id, ";", i)
// 		time.Sleep(time.Millisecond * 1000)
// 	}
// }

// go channels allows us to pass data between go routines

var pizzaNum = 0
var pizzaName = ""

func makeDough(stringChan chan string) {
	pizzaNum++
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)
	fmt.Println("Make Dough and Send for Sauce")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func addSauce(stringChan chan string) {
	pizza := <-stringChan
	fmt.Println("Add Sauce and Send", pizza, "for toppings")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func addToppings(stringChan chan string) {
	pizza := <-stringChan
	fmt.Println("Add Toppings to", pizza, "and ship")

	time.Sleep(time.Millisecond * 10)
}

func main() {

	stringChan := make(chan string)
	for i := 0; i < 3; i++ {
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)

		time.Sleep(time.Millisecond * 5000)
	}

	// go routines
	// basically allow us to run multiple different functions parallel
	// for i := 0; i < 10; i++ {
	// 	go count(i)
	// }
	// time.Sleep(time.Millisecond * 11000)

	// http.HandleFunc("/", handler)
	// http.HandleFunc("/earth", handler2)

	// http.ListenAndServe(":8080", nil)

	// randInt := 5
	// randFloat := 10.5
	// randString := "100"
	// randString2 := "250.5"

	// fmt.Println(float64(randInt))
	// fmt.Println(int(randFloat))

	// newInt, _ := strconv.ParseInt(randString, 0, 64)
	// fmt.Println(newInt)

	// newFloat, _ := strconv.ParseFloat(randString2, 64)
	// fmt.Println(newFloat)

	// file, err := os.Create("samp.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// file.WriteString("This is some random text")
	// file.Close()

	// stream, err := ioutil.ReadFile("samp.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// readString := string(stream)
	// fmt.Println(readString)

	// csvString := "1,2,4,5,6"
	// fmt.Println(strings.Split(csvString, ","))
	// listOfLetters := []string{"c","a","b"}
	// sort.Strings(listOfLetters)
	// fmt.Println("Letters:", listOfLetters)
	// listOfNums := strings.Join([]string{"3","2","1"}, ", ")
	// fmt.Println(listOfNums)

	// sampString := "hello, world"
	// fmt.Println(strings.Contains(sampString, "lo"))
	// fmt.Println(strings.Index(sampString, "lo"))
	// fmt.Println(strings.Count(sampString, "l"))
	// fmt.Println(strings.Replace(sampString, "l", "x", 3))

	// rect := Rectangle{20,50}
	// circ := Circle{4}

	// fmt.Println("Rectangle Area =", getArea(rect))
	// fmt.Println("Circle Area =", getArea(circ))

	// strucks are going to provide us with a way
	// to define our own datatypes
	// rect1 := Rectangle{leftX: 0, topY:50, height: 10, width: 10}
	// or
	// rect1 := Rectangle{0, 50, 10, 10}
	// fmt.Println("Area of rectangle =", rect1.area())
	// fmt.Println("Rectangle is", rect1.width, "wide")

	// fmt.Println(", World!")

	// yPtr := new(int)
	// changeYValNow(yPtr)
	// fmt.Println("y =", *yPtr)

	// x := 0
	// here we pass reference of the value, not the value itself
	// that is the power that we get with pointers
	// we actually allowed to pass the actual memory address in
	// and the assign a value to that memory address rather than
	// passing by values ...
	// changeXValNow(&x)
	// and here we pass value, not the reference to the value
	// changeXVal(x)
	// fmt.Println("x =", x)
	// fmt.Println("Memory Address for x =", &x)

	// it's a sort of a way of using exceptions
	// and catching exeptions in other programming languages
	// demPanic()

	// fmt.Println(safeDiv(3,0))
	// fmt.Println(safeDiv(3,3))

	// defer printTwo()
	// printOne()

	// fmt.Println(factorial(3))

	// closure
	// num3 := 3
	// doubleNum := func() int{
	// 	num3 *= 2
	// 	return num3
	// }
	// fmt.Println(doubleNum())
	// fmt.Println(doubleNum())

	// fmt.Println(subtractThem(1,2,3,4,5))

	// num1, num2 := next2Values(5)
	// fmt.Println(num1, num2)

	// listNumbers := []float64 {1,2,3,4,5}

	// fmt.Println("Sum:", addThemUp(listNumbers))

	// ********* maps - collection of key-value pairs

	// presAge := make(map[string] int)

	// presAge["TheodoreRoosevelt"] = 26
	// fmt.Println(presAge["TheodoreRoosevelt"])
	// fmt.Println(len(presAge))

	// presAge["John F. Kennedy"] = 35
	// fmt.Println(len(presAge))

	// delete(presAge, "TheodoreRoosevelt")
	// fmt.Println(len(presAge))

	// ********* logical operaitor // start

	// slice is like an array but whenever you define them you leave out the size
	// numSlice := []int {5,4,3,2,1}
	// slice of array type with initial value - zero, length - 10
	// numSlice3 := make([]int, 5, 10);
	// fmt.Println(len(numSlice3))
	// fmt.Println(numSlice3)
	// copy(numSlice3, numSlice)
	// fmt.Println(numSlice3)
	// numSlice3 = append(numSlice3, 0, -1)
	// fmt.Println(numSlice3)
	// fmt.Println(len(numSlice3))
	// fmt.Println(numSlice3)
	// now we will actually be able make slice by slicing from another slice
	// numSlice2 := numSlice[:2]
	// fmt.Println("numSlice2 =", numSlice2)
	// fmt.Println("numSlice[2:] =", numSlice[2:])
	// fmt.Println("numSlice[:2] =", numSlice[:2])
	// numSlice2 := numSlice[3:5]
	// fmt.Println("numSlice2[0] =", numSlice2[0])
	// fmt.Println("numSlice2[0] =", numSlice2[1])
	// fmt.Println(len(numSlice2))

	// relational operators

	// var favNums2[5] float64

	// favNums2[0] = 153;
	// favNums2[1] = 3;
	// favNums2[2] = 53;
	// favNums2[3] = 13;
	// favNums2[4] = 15;

	// fmt.Println(favNums2)
	// fmt.Println(favNums2[3])

	// favNums3 := [5]float64 {1,2,3,4,5}
	// fmt.Println(favNums3)

	// for i, value := range favNums3 {
	// 	fmt.Println(value, i)
	// }

	// for _, value := range favNums3 {
	// 	fmt.Println(value)
	// }

	// yourAge := 30

	// switch yourAge {
	// 	case 16 : fmt.Println("Go Drive")
	// 	case 18 : fmt.Println("Go Vote")
	// 	default : fmt.Println("Go Have Fun")
	// }

	// if yourAge >= 16 {
	// 	fmt.Println("you can drive")
	// } else if yourAge >= 18 {
	// 	fmt.Println("you can vote")
	// } else  {
	// 	fmt.Println("you can't drive")
	// }

	// for j := 0; j < 5; j++ {
	// 	fmt.Println(j)
	// }

	// logical operators

	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// fmt.Println("true && false =", true && false)
	// fmt.Println("true || false =", true || false)
	// fmt.Println("!true =", !true)

	// ********* logical operaitor // end

	// ********* println, printf // primitive datatypes // start

	// const pi float64 = 3.14159265
	// var isOver40 bool = false
	// fmt.Printf("%.3f \n", pi);
	// print what datatype is
	// fmt.Printf("%T \n", pi);
	// print what datatype is for bool
	// fmt.Printf("%t \n", isOver40);
	// for integer show up
	// fmt.Printf("%d \n", 100);
	// for binary representation show up
	// fmt.Printf("%b \n", 100);
	// for character code of...
	// fmt.Printf("%c \n", 44);
	// for hex code
	// fmt.Printf("%x \n", 17);
	// for scientific notation
	// fmt.Printf("%e \n", pi);

	// const pi float64 = 3.14159265
	// var isOver40 bool = true

	// fmt.Printf("%.3f \n", pi);

	// var myName string = "ivavovodetstvo"
	// fmt.Println("i like \n\n")
	// fmt.Println("newlines")
	// fmt.Println(" is a robot")
	// fmt.Println(myName + " is a robot")
	// fmt.Println(len(myName))

	// var (
	// 	varA = 2
	// 	varB = 3
	// )

	// fmt.Println(varA, varB, pi)

	// var age int = 40;
	// var favNum float64 = 1.6180339
	// randNum := 1

	// var numOne = 1.000
	// var num99 = .9999

	// fmt.Println("6 + 4 = ", 6 + 4)
	// fmt.Println("6 % 4 = ", 6 % 4)

	// your floats are not always going to be 100% accurate
	// fmt.Println(numOne - num99)
	// % - modulus - which is going to return the reminder of a division
	// fmt.Println(age, favNum)
	// fmt.Println(randNum)
	// fmt.Println(age, " ", favNum)

	// ********* println, printf // primitive datatypes // end
}

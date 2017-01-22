package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var pizzaNum = 0
var pizzaName = ""

func main() {
	fmt.Println("Hello world")

	var age int = 40

	var favNum float64 = 1.3454231

	fmt.Println(age, " ", favNum)

	var myName string = "Sachin"

	fmt.Println(len(myName))

	fmt.Println("i like \n \n")
	fmt.Println("new")

	const pi float64 = 3.14159265

	fmt.Printf("%.3f \n", pi)

	fmt.Printf("%T \n", pi)

	fmt.Printf("%d \n", 100)
	fmt.Printf("%b \n", 100)
	fmt.Printf("%c \n", 44)
	fmt.Printf("%x \n", 17)
	fmt.Printf("%e \n", pi)

	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	yourAge := 18

	if yourAge >= 16 {
		fmt.Println("You can drive")
	} else {
		fmt.Println("You cant drive")
	}

	switch yourAge {

	case 16:
		fmt.Println("go drive")
	case 18:
		fmt.Println("go vote")

	default:
		fmt.Println("go have fun")
	}

	var favNums [5]float64
	favNums[0] = 3.1414

	fmt.Println(favNums[0])

	favNums3 := [5]float64{1, 2, 3}
	fmt.Println(favNums3[0])

	for i, value := range favNums3 {
		fmt.Println(value, i)
	}

	for value := range favNums3 {
		fmt.Println(value)
	}

	numSlice := []int{5, 4, 3, 2, 1}

	numSlice2 := numSlice[3:5]
	fmt.Println(numSlice2[0])

	fmt.Println("numSlice[:2] = ", numSlice[:2])

	fmt.Println("numSlice[2:] = ", numSlice[2:])

	numSlice3 := make([]int, 5, 10)
	copy(numSlice3, numSlice)
	fmt.Println(numSlice3[4])

	numSlice3 = append(numSlice3, 0, -1)
	fmt.Println(numSlice3[6])

	presAge := make(map[string]int)
	presAge["TheodoreRoosevelt"] = 44

	fmt.Println(presAge["TheodoreRoosevelt"])
	fmt.Println(len(presAge))

	delete(presAge, "TheodoreRoosevelt")
	fmt.Println(len(presAge))

	listNums := []float64{1, 2, 3, 4, 5}
	fmt.Println("Sum :=", add(listNums))

	num1, num2 := next2Values(5)
	fmt.Println("next2Values ", num1, num2)
	fmt.Println("subtract all ", substractAll(1, 2, 3, 4, 5))

	num := 3

	doubleNum := func() int {
		num *= 2
		return num
	}

	fmt.Println(doubleNum())
	fmt.Println(doubleNum())
	fmt.Println(factorial(3))

	defer printTwo()
	printOne()

	fmt.Println(safeDivision(3, 0))
	fmt.Println(safeDivision(3, 2))
	demoPanic()

	x := 0
	fmt.Println("original x ", x)
	fmt.Println("memory address for x = ", &x)
	changeXValNow(&x)
	fmt.Println("changed x", x)

	yPtr := new(int)
	fmt.Println("original y ", *yPtr)
	changeYValNow(yPtr)
	fmt.Println("changed y =", *yPtr)

	rect1 := Rectangle{leftX: 0, topY: 10, height: 20, width: 30}
	rect2 := Rectangle{10, 10, 10, 10}

	fmt.Println("Rectangle 1 is ", rect1.width, "wide")
	fmt.Println("Rectangle 2 is ", rect2.height, "height")

	fmt.Println("Area of rectangle =", rect1.area())

	circle1 := Circle{1}
	square1 := Square{2}

	fmt.Println("Area of circle = ", getArea(circle1))
	fmt.Println("Area of square = ", getArea(square1))

	sampString := "Hello world"
	fmt.Println(sampString)
	fmt.Println(strings.Contains(sampString, "lo"))
	fmt.Println(strings.Index(sampString, "lo"))
	fmt.Println(strings.Count(sampString, "l"))
	fmt.Println(strings.Replace(sampString, "l", "x", 3))

	csvString := "1,2,3,4,5,6"
	fmt.Println(strings.Split(csvString, ","))
	listOfLetters := []string{"c", "a", "b"}
	sort.Strings(listOfLetters)
	fmt.Println("Letters : ", listOfLetters)
	listOfNums := strings.Join([]string{"3", "2", "1"}, ", ")
	fmt.Println(listOfNums)

	file, err := os.Create("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("This is some random text")
	file.Close()
	stream, err := ioutil.ReadFile("sample.txt")

	if err != nil {
		log.Fatal(err)
	}
	readString := string(stream)
	fmt.Println(readString)

	randInt := 5
	randFloat := 10.3
	randString := "100"
	randString2 := "100.4"
	fmt.Println(float64(randInt))
	fmt.Println(int(randFloat))
	newInt, _ := strconv.ParseInt(randString, 0, 64)
	fmt.Println(newInt)

	newFloat, _ := strconv.ParseFloat(randString2, 64)
	fmt.Println(newFloat)

	for i := 0; i < 10; i++ {
		go count(i)
	}
	time.Sleep(time.Millisecond * 11000)

	stringChan := make(chan string)

	for i := 0; i < 3; i++ {
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)
	}

	time.Sleep(time.Millisecond * 5000)

	http.HandleFunc("/", handler)
	http.HandleFunc("/earth", handler2)
	http.ListenAndServe(":8080", nil)

}

func makeDough(stringChan chan string) {
	pizzaNum++

	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)

	fmt.Println("make dough & send sauce")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func addSauce(stringChan chan string) {
	pizza := <-stringChan

	fmt.Println("Add sauce and send ", pizza, "for toppings")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func addToppings(stringChan chan string) {
	pizza := <-stringChan

	fmt.Println("Add toppings to ", pizza, "and ship")

	time.Sleep(time.Millisecond * 10)
}

func count(id int) {
	for i := 0; i < 10; i++ {
		fmt.Println(id, ":", i)
		time.Sleep(time.Millisecond * 1000)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Earth\n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HelloEarth\n")
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	leftX  float64
	topY   float64
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

type Square struct {
	side float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (s Square) area() float64 {
	return s.side * s.side
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func changeYValNow(yPtr *int) {
	*yPtr = 100
}

func changeXValNow(x *int) {
	*x = 2
}

func demoPanic() {

	defer func() {
		fmt.Println(recover())
	}()

	panic("PANIC")
}

func safeDivision(num1, num2 int) int {

	defer func() {
		fmt.Println(recover())
	}()

	solution := num1 / num2
	return solution
}

func printOne() {
	fmt.Println(1)
}

func printTwo() {
	fmt.Println(2)
}

func add(numbers []float64) float64 {
	sum := 0.0

	for _, val := range numbers {
		sum += val
	}

	return sum
}

func next2Values(number int) (int, int) {

	return number + 1, number + 2
}

func substractAll(args ...int) int {
	finalVal := 0

	for _, value := range args {
		finalVal -= value
	}
	return finalVal
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

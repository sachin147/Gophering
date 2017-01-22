


package main


import "fmt"


func plus() func(int) int {
	return func(x int) int {
		return x+1
	}
}

func main() {
	rfunc := plus()
	fmt.Printf("%v",rfunc(1))
}

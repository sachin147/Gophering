


package main

import "fmt"

func ascend(x int, y int) (int,int) {
	if x > y {
		return y,x
	}
	return x,y
}

func main() {
	fmt.Println(ascend(2,5))
	fmt.Println(ascend(5,2))
}

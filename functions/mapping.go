


package main

import "fmt"


func Map(f func(int) int, l []int) []int {
	new := make([]int, len(l))

	for i,j := range new {
		new[i] = f(j)
	}
	return new
}


func main() {
	a := []int{1,2,3,4,5}

	f := func(i int) int {
		return i*i*i
	}

//	m := Map(f,a)

	fmt.Printf("%v",(Map(f,a)))
}

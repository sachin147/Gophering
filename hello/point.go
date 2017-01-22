package main

import "fmt"

type Point struct {
	X, Y int
}

var (
	p = Point{1, 2}
	q = &Point{1, 2}
	r = Point{X : 1}
	s = Point{}
)

func main() {
	fmt.Println(p, q, r, s)
}

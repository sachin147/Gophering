package main

import "fmt"

type User struct {
	Id int
	Name string
	Location string
}

func main() {
	x := new(User)
	y := &User{}
	fmt.Println(*x == *y)
}


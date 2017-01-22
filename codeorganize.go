


package main

import (
	"fmt"
)


const(
	const = "constant"
)


var(
	a = 1
)


//struct
type A struct {
	name string
}

//functions
func newA() string {
	return "A"
}



//methods
func (a *A) greet() string {
	return fmt.Sprintf("%s",a.name)
}




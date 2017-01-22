


package main


import "fmt"

type stack struct {
	i int
	data [10]int
}


func (s *stack) push(n int) {
	s.data[s.i] = n
	s.i++
}


func (s *stack) pop() int {
	s.i--
	return s.data[s.i]
}

func main() {
	s := new(stack) //pointer to stack s
	s.push(10)
	s.push(20)
	s.pop()
	fmt.Printf("stack %v\n",s)
}

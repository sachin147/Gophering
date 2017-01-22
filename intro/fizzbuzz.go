


package main


import "fmt"


func main() {

	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("number %d = FizzBuzz\n",i)
		} else if i%3 == 0 {
			fmt.Printf("number %d = Fizz\n",i)
		} else if i%5 == 0 {
			fmt.Printf("number %d = Buzz\n",i)
		}
	}
}

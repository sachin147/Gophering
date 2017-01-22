


package main


import "fmt"


func main() {
	
	var arr[10]int
	
	for i := 0; i < 10; i++ {
		fmt.Printf("looping at %d\n",i)
		arr[i] = i
		fmt.Println(arr)
	}
	fmt.Printf("%v",arr)
}

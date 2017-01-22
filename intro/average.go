

package main


import "fmt"


func main() {

	nums := make([]float64, 10)

	sum := 0.0

	for i:=0; i<len(nums); i++ {
		nums[i] = float64(i)
		sum+=float64(i)
	}
	fmt.Println("%v",nums)
	fmt.Println(sum/float64(len(nums)))
}

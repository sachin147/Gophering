

package main


import "fmt"


func main() {
	var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt","Emma", "Isabella", "Emily", "Madison","Ava", "Olivia", "Sophia", "Abigail","Elizabeth", "Chloe", "Samantha","Addison", "Natalie", "Mia", "Alexis"}


	var maxLength int

	for _, name := range names {
		if l := len(name); l > maxLength {
			maxLength = l  //set  size of 2D slice
		}
	}

	
	output := make([][] string, maxLength) // create 2D slice of strings with given size


	for _,name := range names {
		output[len(name) - 1] = append(output[len(name) -1], name) // populate 2D slice , append appends name to existing slice at position len(name) -1 of output 2D slice
	}

	fmt.Printf("%v", output)
}

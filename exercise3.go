


package main


import "strings"
import "fmt"

func WordCount(s string) map[string] int {

	words := strings.Fields(s)
	count := make(map[string]int)

	for _, word := range words {
		count[word]++
	}

	return count
}


func main() {
	fmt.Println(WordCount("Sachin T"))
}




package main


import "fmt"
import "unicode/utf8"


func main() {

	count := 0

	line := "hands on go"

	for i:=0; i<len(line); i++ {
		if string(line[i]) != " " {
			count = count + 1
		}
	}
	fmt.Printf("number of chars %d", count)

	fmt.Printf("bytelength %d",len([]byte(line)))
	fmt.Printf("runecount %d", utf8.RuneCount([]byte(line)))
				
}

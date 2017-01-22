

package main


import "fmt"


func main() {
	s := "reverse"
	r := make([]byte,len(s))
	/*DOES NOT WORK
       	for i,j := 0, len(s)-1; i<j; i,j = i+1,j-1 {
		s[i],s[j] = s[j],s[i]	
	}
	*/

	for i:=0; i<len(s); i++ {
		r[i] = s[len(s)-1-i]
	}
	fmt.Println(string(r))

}

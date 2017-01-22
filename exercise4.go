
package main


import "fmt"


var(
	coins = 50
	users = []string{"Matthew", "Emilie", "Augustus"}
	distribution = make(map[string]int, len(users))

)


func getCoinsForUser(user string) int {
	var total int;
	
	for i := 0; i < len(user); i++ {
		switch string(user[i]) {  //cast to string
			case "a", "A": total = total + 1
			case "e", "E": total = total + 1
			case "i", "I": total = total + 2
			case "o", "O": total = total + 3
			case "u", "U": total = total + 4		
		}
	}
	return total
}


func main() {

	for _, user  := range users {
		userCoins := getCoinsForUser(user)
		if userCoins > 10 {
			userCoins = 10
		}
		distribution[user] = userCoins
		coins = coins - userCoins
	}

	fmt.Println(distribution)
	fmt.Println("coins left", coins)
}

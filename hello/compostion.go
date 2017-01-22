package main

import "fmt"


type User struct {
	Id int
	Name, Location string
}

type Player struct {
	User
	GameId int
}


func main() {
	p:= Player {
		User{Id : 42, Name: "Matt", Location: "LA"},
		90404,
	}

	fmt.Printf("Id: %d, Name : %s , Location : %s, Gameid: %d\n",p.Id, p.Name, p.Location, p.GameId)
	p.Id = 11
	fmt.Printf("%+v", p)
}

package main

import "fmt"

type User struct {

	Id int
	Name, Location string

}

func (u *User) Greetings() string {
	return fmt.Sprintf("hi %s from %s", u.Name, u.Location)
}

type Player struct {

	*User
	GameId int
}


func main() {

	p := &Player{ User : &User{Id : 1, Name : "Sachin", Location : "MUM"}, GameId : 1}
	fmt.Println(p.Greetings())
}

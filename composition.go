package main

import "fmt"

type User struct {
	id int
	name string
}


type Player struct {
	User
	gameId int
}

func main() {
	var player Player = Player{User{ 1, "Donuts"}, 2}

	fmt.Println("player", player)
}

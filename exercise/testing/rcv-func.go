package main

import "fmt"

type name string

type player struct {
	Health string
	Energy string
	Name   name
}

func (player *player) increaseHealthAndEnergy(increaseBy string) {
	player.Health = increaseBy
	player.Energy = increaseBy
}

func (player *player) decreaseHealthAndEnergy(decreaseBy string) {
	player.Health = decreaseBy
	player.Energy = decreaseBy
}
func main() {
	player1 := player{Name: "pranav", Health: "medium", Energy: "medium"}
	fmt.Println("Initial:", player1)

	player1.increaseHealthAndEnergy("Max")
	fmt.Println("\nAfter Increase:", player1)

	player1.decreaseHealthAndEnergy("Min")
	fmt.Println("\nAfter Decrease:", player1)

}

//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

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

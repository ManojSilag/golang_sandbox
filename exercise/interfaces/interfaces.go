//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Vehicles interface {
	Lift()
}

type Vehicle struct {
	name, model string
}

func (v Vehicle) Lift() {
	switch v.name {
	case "Motorcycles":
		fmt.Printf("--%v small lifts\n", v.name)
	case "Cars":
		fmt.Printf("--%v standard lifts\n", v.name)
	case "Trucks":
		fmt.Printf("--%v large lifts\n", v.name)
	default:
		panic("Unidentified Vehicle")
	}
}

func liftVehicles(wheelers []Vehicles) {
	for i := 0; i < len(wheelers); i++ {
		v := wheelers[i]
		fmt.Println("--Vehile lifting----")
		v.Lift()

	}
}

func main() {
	wheelers := []Vehicles{Vehicle{name: "Motorcycles", model: "xox"},
		Vehicle{name: "Cars", model: "xox"}}
	liftVehicles(wheelers)
}

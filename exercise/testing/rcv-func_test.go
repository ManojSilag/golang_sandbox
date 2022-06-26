//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

func TestPlayerCreate(t *testing.T) {

	pranav := player{Name: "pranav", Health: "medium", Energy: "medium"}

	if pranav.Name != "pranav" {
		t.Errorf("Name mismatch")
	}

	if pranav.Health != "medium" {
		t.Errorf("Health mismatch")
	}

	if pranav.Energy != "medium" {
		t.Errorf("Energy mismatch")
	}

}

func TestPowerIncrease(t *testing.T) {
	pratik := player{Name: "pratik", Health: "medium", Energy: "medium"}

	pratik.increaseHealthAndEnergy("max")

	if pratik.Health != "max" {
		t.Errorf("Health mismatch after increase")
	}

	if pratik.Energy != "max" {
		t.Errorf("Energy mismatch after increase")
	}
}

func TestPowerDecrease(t *testing.T) {
	pratik := player{Name: "pratik", Health: "medium", Energy: "medium"}

	pratik.decreaseHealthAndEnergy("min")

	if pratik.Health != "min" {
		t.Errorf("Health mismatch after increase")
	}

	if pratik.Energy != "min" {
		t.Errorf("Energy mismatch after increase")
	}
}

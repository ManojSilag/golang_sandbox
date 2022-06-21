package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {

	switch p := price(); {
	case p < 2:
		fmt.Println("Cheap")
		// fallthrough
	case p < 10:
		fmt.Println("Modertely")
	default:
		fmt.Println("Expensice")
	}
}

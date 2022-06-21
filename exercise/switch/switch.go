//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

func main() {

	// age := 2

	switch age := 18; {
	case age == 0:
		fmt.Println("New Born")
	case age >= 1 && age <= 3:
		fmt.Println("Toddler")
	case age >= 4 && age <= 12:
		fmt.Println("Child")
	case 13 <= age && age <= 17:
		fmt.Println("Teenager")
	default:
		fmt.Println("adult")
	}
}

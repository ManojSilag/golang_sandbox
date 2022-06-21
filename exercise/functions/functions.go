//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greeting(name string) {
	fmt.Println("Hell0", name, "there !")
}

func message() string {
	return "how are you ?"
}

func addThreeNum(a, b, c int) int {
	return a + b + c
}

func number() int {
	return 12
}

func numbers() (int, int) {
	return 15, 15
}

func main() {
	greeting("Manoj")

	fmt.Println(message())

	sum := addThreeNum(1, 2, 3)
	fmt.Println("sum of three is", sum)

	firstNum := number()

	secondNum, thirdNum := numbers()

	summ := addThreeNum(firstNum, secondNum, thirdNum)
	fmt.Println("sum of three number is", summ)

}

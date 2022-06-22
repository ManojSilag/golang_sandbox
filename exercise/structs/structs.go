//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
)

type rectangle struct {
	length int
	width  int
}

func area(rec rectangle) int {
	return rec.length * rec.width
}

func perimeter(rec rectangle) int {
	return (2 * rec.length) + (2 * rec.width)
}
func main() {

	rec1 := rectangle{length: 15, width: 5}
	fmt.Println(rec1)

	aarea := area(rec1)
	fmt.Println("area", aarea)

	pperimiter := perimeter(rec1)
	fmt.Println("perimiter", pperimiter)

	rec1.length = rec1.length + rec1.length
	rec1.width = rec1.width + rec1.width

	fmt.Println(rec1)

	aarea = area(rec1)
	fmt.Println("area", aarea)

	pperimiter = perimeter(rec1)
	fmt.Println("perimiter", pperimiter)

}

//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func PrintAssemblyLine(line []Part) {
	fmt.Println()

	for j := 0; j < len(line); j++ {
		part := line[j]
		fmt.Println("Part", part)
	}
}
func main() {
	var assemblyLine []Part
	PrintAssemblyLine(assemblyLine)

	assemblyLine = append(assemblyLine, "partOne", "partTwo", "partThree")
	PrintAssemblyLine(assemblyLine)

	assemblyLine = append(assemblyLine, "partFour", "partFive")
	PrintAssemblyLine(assemblyLine)

	assemblyLine = assemblyLine[3:]
	PrintAssemblyLine(assemblyLine)

}

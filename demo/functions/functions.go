package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello from greet")
}

func main() {
	greet()

	doubled := double(5)
	fmt.Println(doubled)

	sum := add(doubled, 4)
	fmt.Println(sum)

	summed := add(doubled, sum)
	fmt.Println(summed)
}

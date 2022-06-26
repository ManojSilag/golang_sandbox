package main

import "fmt"

func sum(nums ...int) int {
	sum := 0

	for _, n := range nums {
		sum += n
	}

	return sum
}

func main() {
	fmt.Println(sum(2 + 3))
	fmt.Println(sum(2, 0))
	fmt.Println(sum(2, 3, 5))

}

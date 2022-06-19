package main

import "fmt"

func main() {
	var myname = "manoj"
	fmt.Println("my name is", myname)

	var name string = "kathy"
	fmt.Println("name =", name)

	username := "admin"
	fmt.Println("username = ", username)

	var sum int
	fmt.Println("The sum is", sum)

	part1, other := 1, 5
	fmt.Println("part 1 is", part1, "other is", other)

	part2, other := 2, 0
	fmt.Println("part 2 is", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("The sum is", sum)

	var (
		lessonName = "variables"
		lessontype = "Demo"
	)

	fmt.Println("lessonName =", lessonName)
	fmt.Println("lessontype =", lessontype)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)

}

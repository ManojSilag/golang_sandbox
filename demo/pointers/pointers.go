package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1
	fmt.Println("counter", counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {
	str := "manoj"
	var strPtr *string = &str

	counter := Counter{hits: 0}

	replace(strPtr, "Manoj", &counter)
	fmt.Println(counter)
	fmt.Println(*strPtr)
}

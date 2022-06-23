package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)

	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println(shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	// shoppingList["cereal"] = 23
	cereal, found := shoppingList["cereal"]
	if !found {
		fmt.Println("ceral no need")
	} else {
		fmt.Println(cereal, "cereal needed")
	}

	totalItems := 0

	for _, value := range shoppingList {
		totalItems += value
	}

	fmt.Println("TotalItems", totalItems)
}

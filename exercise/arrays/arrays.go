//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type product struct {
	name  string
	price int
}

func statistics(itemss [4]product) {

	var lastitem product
	var totalitem int
	var totalcost int
	for i := 0; i < len(itemss); i++ {
		// fmt.Println(itemss[i])
		item := itemss[i]
		totalcost += item.price
		if item.name != "" && item.price != 0 {
			totalitem += 1
			lastitem = item
		}
	}
	fmt.Println("Total cost", totalcost)
	fmt.Println("Total items", totalitem)
	fmt.Println("Lastitem", lastitem)

}

func main() {

	items := [4]product{
		{name: "choclate", price: 10},
		{name: "Ice cream", price: 25},
		{name: "waffels", price: 180},
	}

	statistics(items)

	items[3] = product{name: "food", price: 400}

	statistics(items)

}

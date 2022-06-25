//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

//* Create a structure to store items and their security tag state
type securityTag struct {
	status bool
	item   string
}

//  - Security tags have two states: active (true) and inactive (false)
const (
	active   = true
	inactive = false
)

func statusChange(statusPtr *securityTag) {
	statusPtr.status = !statusPtr.status
}

func checkout(tags []securityTag) {
	for i := 0; i < len(tags); i++ {
		statusChange(&tags[i])
	}
}

func main() {
	//  - Create at least 4 items, all with active security tags
	tag1 := securityTag{status: active, item: "bag"}
	tag2 := securityTag{status: active, item: "pen"}
	tag3 := securityTag{status: active, item: "bottle"}
	tag4 := securityTag{status: active, item: "compass"}

	//  - Store them in a slice or array
	tags := []securityTag{tag1, tag2, tag3, tag4}
	fmt.Println(tags)

	//  - Deactivate any one security tag in the array/slice
	statusChange(&tags[2])
	fmt.Println(tags)

	checkout(tags)
	fmt.Println(tags)
}

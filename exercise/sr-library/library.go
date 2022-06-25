//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Title string
type Name string

type LendAudit struct {
	checkout time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int
	lended int
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkout.String(), "through", returnTime)
	}
}

func printMemberAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}
	fmt.Println()
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of Library")
		return false
	}

	if book.lended == book.total {
		fmt.Println("No more books available to lend")
		return false
	}

	book.lended += 1
	library.books[title] = book
	member.books[title] = LendAudit{checkout: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of Library")
		return false
	}

	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did not check out this book")
		return false
	}

	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

func main() {
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["Webapps in Go"] = BookEntry{
		total:  4,
		lended: 0,
	}

	library.books["The little Go back"] = BookEntry{
		total:  4,
		lended: 0,
	}

	library.books["Lets learn GO"] = BookEntry{
		total:  2,
		lended: 0,
	}

	library.books["Go Bootcamp"] = BookEntry{
		total:  1,
		lended: 0,
	}

	library.members["Jayson"] = Member{"Jayson", make(map[Title]LendAudit)}
	library.members["Vijay"] = Member{"Vijay", make(map[Title]LendAudit)}
	library.members["Sam"] = Member{"Sam", make(map[Title]LendAudit)}

	fmt.Println(library)
	fmt.Println("\nInitial:")
	printLibraryBooks(&library)
	printMemberAudits(&library)

	member := library.members["Jayson"]
	checkedOut := checkoutBook(&library, "Go Bootcamp", &member)
	fmt.Println("\nCheck out a book:")
	if checkedOut {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}

	returned := returnBook(&library, "Go Bootcamp", &member)
	fmt.Println("\nCheck in a book:")

	if returned {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}

}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// type book struct {
// 	name     string
// 	checkout time.Time
// 	returned time.Time
// }

// type member struct {
// 	name string
// }

// func library(books []book, members []member) {
// 	fmt.Println("--books--")
// 	for i := 0; i < len(books); i++ {
// 		fmt.Println(books[i].name, "is checked out at", books[i].checkout, "and returned at", books[i].returned)
// 	}

// 	fmt.Println("--member--")
// 	for i := 0; i < len(members); i++ {
// 		fmt.Println(members[i].name)
// 	}
// }

// func main() {
// 	book1 := book{name: "Alice"}
// 	book2 := book{name: "In"}
// 	book3 := book{name: "Wonderland"}
// 	book4 := book{name: "boxx"}

// 	books := []book{book1, book2, book3, book4}

// 	member1 := member{"manoj"}
// 	member2 := member{"silag"}
// 	member3 := member{"kim"}

// 	members := []member{member1, member2, member3}

// 	// fmt.Println(books)
// 	// fmt.Println(members)

// 	library(books, members)

// 	books[1].checkout = time.Now()
// 	library(books, members)

// 	books[1].returned = time.Now()
// 	library(books, members)

// }

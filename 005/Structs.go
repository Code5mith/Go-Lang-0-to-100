package main

import "fmt"

// This struct defines a "Book" with relevant properties
type Book struct {
	Title     string   // Title of the book
	Author    string   // Author of the book
	ISBN      string   // International Standard Book Number (unique identifier)
	Year      int      // Year the book was published
	IsFiction bool     // Whether the book is fiction or non-fiction
	Reviews   []string // Array of user reviews (optional)
}

func main() {
	// Create a new Book instance
	myBook := Book{
		Title:     "The Hitchhiker's Guide to the Galaxy",
		Author:    "Douglas Adams",
		ISBN:      "978-0345391803",
		Year:      1979,
		IsFiction: true,
	}

	// Print book details using a formatted string
	fmt.Printf("Title: %s\nAuthor: %s\nISBN: %s\nYear: %d\nFiction: %t\n",
		myBook.Title, myBook.Author, myBook.ISBN, myBook.Year, myBook.IsFiction)

	// Access and modify individual fields
	myBook.Reviews = append(myBook.Reviews, "A hilarious and thought-provoking read!")

	// Print book details with the added review
	fmt.Println("\nReviews:")
	for _, review := range myBook.Reviews {
		fmt.Println("-", review)
	}
}

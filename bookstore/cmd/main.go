package main

import "fmt"
type Book struct {	//Key Value pairs?
	Title string
	Author string
	SeriesNumber int
	PriceCents int
	Copies int
}

var books []Book

func AddBook(catalog []Book, book Book) []Book {
	catalog = append(catalog, book)
	return catalog
}

func main() {

	
	books = []Book{
		{
			Title: "The Lost World",
			Author: "Michael Crichton",
			SeriesNumber: 2,
			PriceCents: 499,
			Copies: 1,
		},
		{
			Title: "Jurassic World",
			Author: "Michael Crichton",
			SeriesNumber: 1,
			PriceCents: 499,
			Copies: 3,
		},
	}

	books = AddBook(books, Book{Title: "Spark Joy",Author: "Bill Gates", PriceCents: 1699, Copies: 5})

	for _, b := range books {
		fmt.Printf("We have %v copies of %s by %s", b.Copies, b.Title, b.Author)
		fmt.Println()
	}
}

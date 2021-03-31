package main

import "fmt"

func main() {
	title := "The Lost World"
	author := "Michael Crichton"
	edition := 1.5
	inStock := true
	royaltyPerc := .05
	onSpecial := false
	discount := .15
	copies := 99

	fmt.Println("Title: ", title)
	fmt.Println("Author: ", author)
	fmt.Println("Edition: ", edition)
	fmt.Println("In Stock?: ", inStock)
	fmt.Println("Royalty Percentage", royaltyPerc)
	fmt.Println("How many copies?: ", copies)
	fmt.Println("Is it on special?: ", onSpecial)
	fmt.Println("Discount Amount: ", discount)
}

package main

import (
	"fmt"
	carddeck "gosino/internal"
)

func main() {
	fmt.Println("Welcome to the Casino!")

	deck := carddeck.New()

	deck.Print()

	deck.Shuffle()

	deck.Print()
}

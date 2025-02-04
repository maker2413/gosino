package carddeck

import (
	"fmt"
	"math/rand"
)

type Card struct {
	Suit string
	Rank string
}

type Deck []Card

func New() Deck {
	var deck Deck

	ranks := []string{
		"Ace", "Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Jack", "Queen", "King",
	}

	suits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}

	for s := 0; s < len(suits); s++ {
		for r := 0; r < len(ranks); r++ {
			card := Card{
				Suit: suits[s],
				Rank: ranks[r],
			}

			deck = append(deck, card)
		}
	}

	return deck
}

func (d Deck) Print() {
	for _, card := range d {
		fmt.Printf("%s of %s\n", card.Rank, card.Suit)
	}
}

// Shuffle will shuffle the cards in the deck using the Fisher-Yates Shuffle
// algorithm: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
func (d Deck) Shuffle() {
	for i := len(d) - 1; i > 0; i-- {
		r := rand.Intn(i)

		if i != r {
			d[i], d[r] = d[r], d[i]
		}
	}
}

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []card

type card struct {
	suit  string
	value string
}

/* This function creates a new deck from zero */
func newDeck() deck {
	cards := []card{}

	//Declare all suits and existing values
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// For every card suit it iterates every value and creates a new card
	// Then the new card is appended to the cards slice
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, card{
				suit:  suit,
				value: value,
			})
		}
	}

	return cards
}

/* This function deals the cards from the deck to a hand of type deck*/
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/* This function prints an existing deck*/
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

/* This function converts a slice of cards into a single string*/
func toString(s []card) string {
	var result string
	var cards []string
	for _, card := range s {
		cards = append(cards, string(card.suit+" "+card.value))
		result = strings.Join(cards, ",")
	}
	return result
}

/* This function converts a slice of strings into a slice of cards*/
func toDeck(s []string) []card {
	var deck []card
	for _, string := range s {
		c := strings.Split(string, " ")
		deck = append(deck, card{
			suit:  c[0],
			value: c[1],
		})
	}
	return deck
}

/* This function saves an existing deck into the hard drive */
func (d deck) saveCards(filename string) error {
	return os.WriteFile(filename, []byte(toString(d)), 0666)
}

/* This function reads an existing deck from the hard drive */
func readCardsFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	//If an error occurs it logs the error
	if err != nil {
		log.Fatal("Error: ", err)
	}
	//Splits a single string into a slice of strings
	s := strings.Split(string(bs), ",")

	// Iterates the slice of strings to populate the deck and returns it
	return toDeck(s)
}

/* This function shuffles an existing deck */
func (d deck) shuffle() {

	//A new seed is created by using the actual moment and then converting it into type int64
	src := rand.NewSource(time.Now().UnixNano())
	//The seed is used as a source to randomize the deck's order
	r := rand.New(src)

	//For every card in the deck it creates a random number between the first and the last position
	for i := range d {
		ri := r.Intn(len(d) - 1)
		//The card existing in the actual index is swapped with another in a random position
		d[i], d[ri] = d[ri], d[i]
	}
}

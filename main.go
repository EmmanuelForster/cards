package main

func main() {
	deck := newDeck()

	_, remainingCards := deal(deck, 3)

	remainingCards.saveCards("remaining_cards")

	rc := readCardsFromFile("remaining_cards")

	rc.shuffle()

	rc.print()
}

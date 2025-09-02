package main

func main() {
	cards := newDeck()
	// cards.print()
	cards.saveToFile("deck")
	// readFromFile().print()
	cards.shuffle()
	hand, rest := deal(cards, 5)
	hand.print()
	rest.print()
}

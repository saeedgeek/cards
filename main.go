package main

func main() {
	cards := newDeck()
	cards.print()

	cards = cards.shuffle()
	cards.print()
}

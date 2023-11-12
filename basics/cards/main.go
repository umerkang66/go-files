package main

func main() {
	// this functions are coming from 'deck.go' file
	cards := newDeck()
	hands, remainingCards := deal(cards, 5)
	
	hands.print()
	remainingCards.print()
}

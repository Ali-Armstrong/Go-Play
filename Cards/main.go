package main

func main(){
	//var cards string = "House of spades"
	//cards := newDeckOfCards()
	//cards.saveToFile("my_cards")
	// pick, remaining := deal(cards, 4)
	// pick.print()
	// remaining.print()

	cards := readFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
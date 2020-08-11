package main

import "fmt"

func main(){
	//cards := newDeck()
	//hand, remainingDeck := deal(cards, 5)
	//hand.print()
	//remainingDeck.print()


	//cards := newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("cards_1")

	cards := newDeckFromFile("cards_1")
	cards.print()
	cards.shuffle()
	fmt.Println()
	cards.print()
}

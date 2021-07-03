package main

func main() {

	cards := newDeck()

	/********function with  multiple return values
	firstsec, remaining_decks := deal(cards, 5)
	firstsec.print()
	remaining_decks.print()*/

	//fmt.Println(cards.toString())
	//err := cards.saveToFile("golearn.txt")
	//if err != nil {
	//	fmt.Println(err)
	//}

	//newCards := readDeckFromFile("golearn.txt")
	//newCards.print()

	cards.shuffleCards()
	cards.print()

}

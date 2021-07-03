package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
//which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//****receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//*****multiple return values
func deal(d deck, num int) (deck, deck) {
	return d[:num], d[num:]
}

//***type conversion : []string(value)
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//***Saving data to the hard drive : file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	ss := strings.Split(string(bs), ",")
	return deck(ss)
}

func (d deck) shuffleCards() {
	t := time.Now()
	source := rand.NewSource(t.UnixNano())
	r := rand.New(source)

	for i := range d {
		randomnum := r.Intn(len(d) - 1)
		d[i], d[randomnum] = d[randomnum], d[i]
	}
}

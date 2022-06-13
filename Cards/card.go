package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string


func(d deck) print(){
	for i, card := range d{
		fmt.Println(i, card)
	}
}

func newDeckOfCards() deck{
	deckOfCards := deck {}

	cardSuits := []string {"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string {"Ace", "Two", "Three"}

	for _, suit := range cardSuits{
		for _, value := range cardValues{
			deckOfCards = append(deckOfCards, value + " of " + suit)
		}
	}

	return deckOfCards
}

func deal(d deck, size int) (deck, deck){
	return d[:size], d[size:]
}

func (d deck) toString() string{
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error{
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666);
}

func readFromFile(filename string) deck{
	bs, err := ioutil.ReadFile(filename);

	if err != nil{
		fmt.Println("Error: ",err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d{
		newPosition := r.Intn(len(d) -1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)
//Create a new type of Deck
//Which is a slice of string

type deck []string

func newDeck() deck {
	cards := deck{}

	//Build a new deck
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
        cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
	    for _, value := range cardValues {
	        cards = append(cards, value+" of "+suit)
	    }
	}

	//Return a new deck
	return cards
}


func (d deck) print(){
    for i, card := range d {
        fmt.Println(i, card)
    }
}

func deal(d deck, handSize int) (deck, deck){
    fmt.Println(d)
    return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
    return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
    return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string, ) deck {
    bs, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(err)
	os.Exit(1)
    }
    return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle () {
    length := len(d)
    nowTimestamp := time.Now().UnixNano()
    newRand := rand.New(rand.NewSource(nowTimestamp))

    for i := range d {
	    newPosition := newRand.Intn(length - 1)
	    d[i], d[newPosition] = d[newPosition], d[i]
    }
}

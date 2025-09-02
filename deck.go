package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(name string) error {
	stringed := d.toString()
	message := []byte(stringed)
	err := ioutil.WriteFile(name, message, 0644)
	return err
}

func readFromFile(name string) deck {
	content, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	stringDeck := string(content)
	actualDeck := strings.Split(stringDeck, ",")
	return actualDeck
}

func (d deck) shuffle() deck {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
	return d
}

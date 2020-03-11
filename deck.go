package main

import (
	"fmt"
	"math/rand"
)

type deck []string

func newDeck() deck {
	emtiaz := []string{
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"10",
		"sarbaz",
		"bibi",
		"shah",
		"A",
	}
	khal := []string{
		"pic",
		"del",
		"khesht",
		"gish",
	}

	d := deck{}
	for _, k := range khal {
		for _, e := range emtiaz {
			l := e + " " + k
			d = append(d, l)
		}

	}
	return d

}

func (d deck) print() {
	for i, c := range d {
		fmt.Println(i, c)
	}

}

func (d deck) shuffle() deck {
	newDeck := deck{}
	for {
		lend := len(d)
		if lend <= 0 {
			break
		}
		ran := rand.Intn(lend)

		newDeck = append(newDeck, d[ran])
		d = append(d[:ran], d[ran+1:]...)

	}
	return newDeck

}

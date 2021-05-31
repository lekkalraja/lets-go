package main

import (
	"os"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	length := 56
	firstCard := "Ace of Spades"
	lastCard := "King of Clubs"

	d := newDeck()

	if len(d) != length {
		t.Errorf("Expected : %v but got : %v", length, len(d))
	}

	if d[0] != firstCard {
		t.Errorf("Expected : %v but got : %v", firstCard, d[0])
	}

	if d[len(d)-1] != lastCard {
		t.Errorf("Expected : %v but got : %v", lastCard, d[len(d)-1])
	}
}

func TestDeal(t *testing.T) {
	d := deck{"one", "two", "three", "four"}

	first, last := deal(d, 2)

	if len(first) != 2 || len(last) != 2 {
		t.Errorf("Expected : %v,%v but got : %v,%v", 2, 2, len(first), len(last))
	}
}

func TestToString(t *testing.T) {
	d := deck{"one", "two", "three", "four"}

	ds := d.toString()
	es := "one,two,three,four"

	if strings.Compare(ds, es) != 0 {
		t.Errorf("Expected : %v but got : %v", ds, es)
	}
}

func TestShuffle(t *testing.T) {
	d := deck{"one", "two", "three", "four"}

	d.shuffle()

	if len(d) != 4 {
		t.Errorf("Expected : %v but got : %v", 4, len(d))
	}
}

func TestSaveToFileAndReadFromFile(t *testing.T) {
	fileName := "_decktestfile"

	os.Remove(fileName)

	d := deck{"one", "two", "three", "four"}
	d.saveToFile(fileName)

	loadedDeck, _ := readFromFile(fileName)

	if len(loadedDeck) != 4 {
		t.Errorf("Expected : %v but got : %v", 4, len(loadedDeck))
	}
	os.Remove(fileName)
}

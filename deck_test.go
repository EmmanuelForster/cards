package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0].suit != "Spades" && d[0].value != "Ace"{
		t.Errorf("Expected first card Ace of Spades, but got %+v", d[0])
	}

	if d[len(d)-1].suit != "Clubs" && d[len(d)-1].suit != "Four" {
		t.Errorf("Expected last card Four of Clubs, but got %+v", d[len(d)-1])
	}

}

func TestSaveCardsAndReadCardsFromFile (t *testing.T){
	os.Remove("_decktesting")

	d:=newDeck()
	d.saveCards("_decktesting")

	cff:= readCardsFromFile("_decktesting")

	if len(cff) != 16{
		t.Errorf("Expected 16 cards, but got %v", len(cff))
	}

	os.Remove("_decktesting")
}
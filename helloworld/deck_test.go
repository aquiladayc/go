package main

import (
	"testing"
        "os"
)
func TestNewDeck(t *testing.T) {
    d := newDeck()
    if len(d) != 16 {
        t.Errorf("Expected deck length of 16, but got %v", len(d))
    }

    if d[0] != "Ace of Spades" {
        t.Errorf("First cards should be Ace of Spades but got %v", d[0])
    }
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
    os.Remove("_decktesting")

    d := newDeck()
    d.saveToFile("_decktesting")
    
    loadedDeck := newDeckFromFile("_decktesting")

    if len(loadedDeck) != 16 {
        t.Errorf("Expected deck length of 16, but got %v", len(d))
    }

    os.Remove("_decktesting")
}

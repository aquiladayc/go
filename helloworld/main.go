package main

func main(){
    cards := newDeckFromFile("test_cards")
    cards.print()
    cards.shuffle()
    cards.print()
}



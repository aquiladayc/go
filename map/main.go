package main

import "fmt"

func main() {
    colors := map[string]string{
         "red": "#ff0000",
	 "green": "#7dkr0",
    }
    colors["white"] = "#ddddd"
    printMap(colors)
}

func printMap(c map[string]string){
    for color, hex := range c{
	fmt.Println("Color:", color, " Hex:", hex)
    }
}

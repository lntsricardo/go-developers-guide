package main

import "fmt"

func main() {
	//first method
	// var colors map[string]string

	//second method
	// colors := make(map[string]string)

	// third method
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#8eu839",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

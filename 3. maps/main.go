package main

import "fmt"

func printMap(c map[string]string) {
	for key, val := range c {
		fmt.Printf("Hex code for %v is %v\n", key, val)
	}
}

func main() {
	// var colors map[string]string // Of course maps can be defined without initialization
	// colors :=make(map[string]string) // make() can be used to initialize variables of type slice, map or channels
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "00ff00",
		"blue":  "0000ff",
		"black": "000000",
	}
	colors["white"] = "ffffff"
	delete(colors, "white")
	printMap(colors)
}

package main

import "fmt"

func main() {
	//var colors map[string]string

	//colors := make(map[string]string)

	colors := map[string]string{
		"red": "#FF0000",
		"green": "#4G1234",
	}

	colors["white"] = "ffffff"

	//delete(colors, "white")

	//fmt.Println(colors)

	printColors(colors)
}

func printColors(c map[string]string){
	// iterate over a map
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

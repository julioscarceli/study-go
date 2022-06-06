package main

// 1 Crie um programa que utilize a declaração switch,
// sem switch statement especificado.

// 2Crie um programa que utilize a declaração switch,
// onde o switch statement seja uma variável do tipo string com
// identificador "esporteFavorito"

import (
	"fmt"
)

func main() {
	shirtsize := "M"
	switch {
	case shirtsize == "M":
		fmt.Println("We have available the colors, blue, red and black")
	case shirtsize == "P":
		fmt.Println("We have available the colors, blue and red.")
	case shirtsize == "G":
		fmt.Println("We have available the colors, blue, red and green.")
	case shirtsize == "GG":
		fmt.Println("We have blue and green colors available.")
	}

	//2
	favoritesport := "surfing"
	switch {
	case favoritesport == "surfing":
		fmt.Println("you need to go to hawaii")
	case favoritesport == "soccer":
		fmt.Println("you need to go to brazil")
	case favoritesport == "basketball":
		fmt.Println("you need to go to Eua")

	}

}

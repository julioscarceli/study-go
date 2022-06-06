package main

import "fmt"

// Põe na tela: todos os números de 1 a 10

func main() {
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}

	// Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto,
	//três vezes cada.
	// Por exemplo: 65 U+0041 'A' U+0041 'A' U+0041 'A' 66 U+0042 'B' U+0042
	//'B' U+0042 'B' ...e por aí vai.

	for y := 65; y <= 90; y++ {
		fmt.Println(y)
		for i := 1; i <= 3; i++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

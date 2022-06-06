package main

import "fmt"

func main() {

	// || -> ou
	x := 5
	if (x == 5) || (x == 3) {
		fmt.Println("chis e igual a 3 ou 5")
	}

	// && -> e
	y := 6
	if (y%2 == 0) && (y%3 == 0) {
		fmt.Println("é multiplo de dois e tambem de tres")
	}

	// || -> ou
	i := 4
	if (i%2 == 0) || (i%3 == 0) {
		fmt.Println("é multiplo de dois ou de tres")
	}

}

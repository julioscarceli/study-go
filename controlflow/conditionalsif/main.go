package main

import "fmt"

func main() {
	//if / else.
	if x := 3; x > 5 {
		fmt.Println("You are approved")
	} else {
		fmt.Println("You are desapproved")
	}

	// if, else if, else if.
	if x := 10; x == 10 {
		fmt.Println("top score")

	} else if x <= 5 {
		fmt.Println("you are desapproved")
	} else if x > 5 {
		fmt.Println("you are approved")
	}

	// if,else if,else if, else

	medindoafebre := 40.2
	if (medindoafebre >= 33) && (medindoafebre <= 35) {
		fmt.Println("Voce não tem febre, volte ao trabalho")
	} else if (medindoafebre > 35) && (medindoafebre <= 37.9) {
		fmt.Println("Voce esta com febre baixa")
	} else if (medindoafebre >= 38) && (medindoafebre <= 39) {
		fmt.Println("Sua febre esta alta, va tomar medicação")
	} else {
		fmt.Println("Febre muito alta, va direto para o hospital")
	}

	if num := 13; num < 0 {
		fmt.Println("is negative")
	} else if num < 10 {
		fmt.Println("has 1 digit")
	} else {
		fmt.Println("has more than 1 digit")
	}
}

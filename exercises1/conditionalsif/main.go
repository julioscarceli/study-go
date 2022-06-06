package main

import "fmt"

//1 Crie um programa que demonstre o funcionamento da declaração if.

// Utilizando a solução anterior, adicione as opções else if e else.

func main() {

	multa := false
	if multa {
		fmt.Println("voce perdeu pontos na carteira")
	} else {
		fmt.Println("Voce não perdeu pontos na carteira")
	}

	gravidadedamulta := "grave"
	if gravidadedamulta == "gravissima" {
		fmt.Println("voce perdeu 7 pontos na carteira")
	} else if gravidadedamulta == "grave" {
		fmt.Println("voce perdeu 5 pontos na carteira")
	} else if gravidadedamulta == "media" {
		fmt.Println("voce perdeu 4 pontos na carteira")
	} else {
		fmt.Println("voce perdeu 3 pontos na carteira")

	}
}

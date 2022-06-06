package main

import "fmt"

//Crie e use um struct anônimo.
//Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.

func main() {

	vendacamisetas := struct {
		nome          string
		sobrenome     string
		cores         []string
		idadeetamanho map[int]string
	}{
		nome:      "julio",
		sobrenome: "scarceli",
		cores:     []string{"azul", "vermelho", "preto"},
		idadeetamanho: map[int]string{
			26: "m"},
	}

	fmt.Println(vendacamisetas)

}

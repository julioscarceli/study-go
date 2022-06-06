package main

//Crie um tipo "pessoa" com tipo subjacente struct,
//que possa conter os seguintes campos:
// Nome
//Sobrenome
//Sabores favoritos de sorvete
//Crie dois valores do tipo "pessoa" e demonstre estes valores,
//utilizando range na slice que contem os sabores de sorvete.

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	pessoa1 := pessoa{
		nome:      "Julio",
		sobrenome: "Scarceli",
		sabores:   []string{"Limão", "chocolate", "uva", "morango"}}

	pessoa2 := pessoa{
		nome:      "carol",
		sobrenome: "prates",
		sabores:   []string{"Laranja", "goiaba", "coco", "chocolate"}}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, v := range pessoa1.sabores {
		fmt.Println("\t", v)
	}

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, v := range pessoa2.sabores {
		fmt.Println("\t", v)
	}
}

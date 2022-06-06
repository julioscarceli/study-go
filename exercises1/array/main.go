package main

import "fmt"

//Usando uma literal composta:
//Crie um array que suporte 5 valores do tipo int
//Atribua valores aos seus índices
//Utilize range e demonstre os valores do array.
//Utilizando format printing, demonstre o tipo do array.

func main() {

	array := [3]int{1, 2, 3}
	for i, v := range array {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", array)

}

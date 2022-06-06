package main

import "fmt"

//Crie uma função que receba um parâmetro variádico do tipo int e
//retorne a soma de todos os ints recebidos;
//Passe um valor do tipo slice of int como argumento para a função;
//Crie outra função, esta deve receber um valor do tipo slice of int e
//retornar a soma de todos os elementos da slice;
//Passe um valor do tipo slice of int como argumento para a função.

func main() {
	si := []int{10, 10, 10, 10, 10}
	fmt.Println(funcao1(si...))

	ss := []int{10, 10, 10, 10, 10, 10}
	fmt.Println(funcao2(ss))
}
func funcao1(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func funcao2(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

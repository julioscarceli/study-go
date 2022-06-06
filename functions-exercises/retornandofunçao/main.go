package main

//Crie uma função que retorna uma função.
//Atribua a função retornada a uma variável.
//Chame a função retornada.

import "fmt"

func main() {

	x := 50
	fmt.Println(retornofunc()(x))
}
func retornofunc() func(int) int {
	return func(i int) int {
		return i + 10

	}
}

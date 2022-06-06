package main

import "fmt"

//Pode-se usar uma função como retorno de uma função

func main() {
	x := 5
	fmt.Println(retornaumafunçao()(x))
}
func retornaumafunçao() func(int) int {
	return func(i int) int {
		return i * 10
	}

}

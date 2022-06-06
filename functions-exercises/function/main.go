package main

import "fmt"

//Exercício:
//Crie uma função que retorne um int
//Crie outra função que retorne um int e uma string
//Chame as duas funções
//Demonstre seus resultados

func main() {
	fmt.Println(retornaint())

}
func retornaint() int {
	return 12
}
func retornaintestring() (int, string) {
	return 12, "Doze"
}

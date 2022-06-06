package main

//func (receiver) identifier(parameters) (returns) { code }
// Abstrair funcionalidade
// Reutilização de código
// Funções são definidas com parâmetros
// Funções são chamadas com argumentos

import "fmt"

func main() {
	valor := soma(10, 10)
	fmt.Println(valor)

	basica()

	argumento("morning")

}

func soma(x, y int) int { // funçao com retorno
	return x + y

}
func basica() { // funçao basica
	fmt.Println("hello, go")
}
func argumento(s string) { // funçao que aceita um argumento
	if s == "morning" {
		fmt.Println("Hello, good morning")
	} else if s == "afternoon" {
		fmt.Println("Hello, good afternoon")
	} else {
		fmt.Println("Hello, good night")
	}

}

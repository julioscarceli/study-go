package main

// Em package-level scope, atribua os seguintes valores às variáveis:
// para "x" atribua 26
// para "y" atribua "Harry potter"
// para "z" atribua true
// Na função main:
// Use fmt.Sprintf para atribuir todos esses valores a uma única variável.
//Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o
//operador curto de declaração.
// Demonstre a variável "s".

import "fmt"

var x int = 26
var y string = "Harry Potter"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\n%v\n%v\n", x, y, z)
	fmt.Print(s)

}

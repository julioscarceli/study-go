package main

import "fmt"

// Em package-level scope, utilizando a palavra-chave var,
//crie uma variável com o identificador "y".
//O tipo desta variável deve ser o tipo subjacente do
//tipo que você criou na variables2.
//Na função main:
//Isto já deve estar feito:
//Demonstre o valor da variável "x"
//Demonstre o tipo da variável "x"
//Atribua 42 à variável "x" utilizando o operador "="
//Demonstre o valor da variável "x"
//Agora faça tambem:
//Utilize conversão para transformar o tipo do valor da variável "x"
//em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
//Demonstre o valor de "y"
//Demonstre o tipo de "y"

type funcionario int

var x int
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 12
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}

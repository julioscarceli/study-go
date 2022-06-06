package main

import "fmt"

//Utilize a declaração defer de maneira que demonstre que sua execução
//só ocorre ao final do contexto ao qual ela pertence.

type soma struct {
	primeironumero int
	segundonumero  int
}

func (s soma) somar() {
	fmt.Println(s.primeironumero + s.segundonumero)
}

type subtracao struct {
	primeironumero int
	segundonumero  int
}

func (s subtracao) subtrair() {
	fmt.Println(s.primeironumero - s.segundonumero)

}

func main() {

	calcularsoma := soma{
		primeironumero: 10,
		segundonumero:  10,
	}
	defer calcularsoma.somar() //defer

	calcularsubtracao := subtracao{
		primeironumero: 20,
		segundonumero:  10,
	}
	calcularsubtracao.subtrair()
}

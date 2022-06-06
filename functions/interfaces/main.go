package main

import "fmt"

// Interface determina o conjunto de metodos que vai implementar ela

//example
//Os tipos profissão1 e profissão2 contem o tipo pessoa
//Cada um tem seu método oibomdia(), e podem dar oi utilizando pessoa.oibomdia()
//Implementam a interface gente
//Ambos podem acessar o função serhumano() que chama o método
//oibomdia() de cada gente
//Tambem podemos no método serhumano() tomar ações diferentes dependendo do
//tipo: switch pessoa.(type) { case profissão1:
//fmt.Println(h.(profissão1).valorquesóexisteemprofissão1) [...] }*

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}
type advogado struct {
	pessoa
	numerodeclientes int
	salario          float64
}
type medico struct {
	pessoa
	especializaçao string
	salario        float64
}

func (x advogado) oibomdia() { //método
	fmt.Println("Meu nome é", x.nome, "e eu tenho", x.numerodeclientes, "clientes, bom dia!")
}

func (x medico) oibomdia() { //método
	fmt.Println("Meu nome é", x.nome, "bom dia!")
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	g.oibomdia()
}

func main() {
	mradvogado := advogado{
		pessoa: pessoa{
			nome:      "Marshall",
			sobrenome: "marshmellow",
			idade:     27,
		},
		numerodeclientes: 6,
		salario:          15.100,
	}
	mrmedico := medico{
		pessoa: pessoa{
			nome:      "ted",
			sobrenome: "mosby",
			idade:     26,
		},
		especializaçao: "cardiologista",
		salario:        12.500,
	}

	mradvogado.oibomdia()
	mrmedico.oibomdia()

	serhumano(mradvogado)
	serhumano(mrmedico)

}

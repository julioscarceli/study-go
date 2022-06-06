package main

import "fmt"

//Crie um tipo struct "pessoa" que contenha os campos:
//nome
//sobrenome
//idade
//Crie um método para "pessoa" que demonstre o nome completo e a idade;
//Crie um valor de tipo "pessoa";
//Utilize o método criado para demonstrar esse valor.

type person struct {
	name     string
	lastname string
	age      int
}

func (p person) dadoscompletos() {
	fmt.Println("Hi, my name is", p.name, p.lastname, "and i have", p.age, "years old")
}
func main() {
	Harry := person{"Harry", "Potter", 14}
	Harry.dadoscompletos()
}

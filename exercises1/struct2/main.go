package main

//Utilizando a solução anterior do exercicio 1 do struct,
//coloque os valores do tipo
//"pessoa" num map, utilizando os sobrenomes como key.
//Demonstre os valores do map utilizando range.
//Os diferentes sabores devem ser demonstrados utilizando outro range,
// dentro do range anterior.

import "fmt"

type person struct {
	name     string
	lastname string
	tastes   []string
}

func main() {
	meumapa := make(map[string]person)

	meumapa["Scarceli"] = person{
		name:     "Julio",
		lastname: "Scarceli",
		tastes:   []string{"Cheese", "Pepperoni", "Mushroom", "Green bell pepper"}}

	meumapa["Prates"] = person{
		name:     "Carol",
		lastname: "Prates",
		tastes:   []string{"Sausage", "Onion", "Olive", "Tomato"}}

	for _, valor := range meumapa {
		fmt.Println("Meu nome é :", valor.name)
		for _, valor := range valor.tastes {
			fmt.Println("-", valor)
		}
	}

}

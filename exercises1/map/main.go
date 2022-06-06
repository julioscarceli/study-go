package main

import "fmt"

// Crie um map com key tipo string e value tipo []string.
// Key deve conter nomes no formato sobrenome_nome
// Value deve conter os hobbies favoritos da pessoa
// Demonstre todos esses valores e seus índices.
func main() {
	mepe := map[string]string{
		"scarceli_julio": "play football",
		"prates_carol":   "drink a beer",
		"scarceli_paula": "working",
		"scarceli_joao":  "sleeping",
	}
	for key, v := range mepe {
		fmt.Println("the user's hobby: ", key, " is:", v)
	}

}

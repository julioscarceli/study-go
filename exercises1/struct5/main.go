package main

import "fmt"

type person struct {
	name     string
	lastname string
	tastes   []string
}

func main() {

	person1 := person{
		name:     "Julio",
		lastname: "Scarceli",
		tastes:   []string{"Cheese", "Pepperoni", "Mushroom", "Green bell pepper"}}

	person2 := person{
		name:     "Carol",
		lastname: "Prates",
		tastes:   []string{"Sausage", "Onion", "Olive", "Tomato"}}

	fmt.Println("\n", person1.name, person1.lastname)
	for _, valor := range person1.tastes {
		fmt.Println("-", valor)
	}

	fmt.Println("\n", person2.name, person2.lastname)
	for _, valor := range person2.tastes {
		fmt.Println("-", valor)
	}

}

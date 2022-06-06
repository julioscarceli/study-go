package main

//Struct is a composite data type that allows us to store values of
//different types.

import "fmt"

type cliente struct {
	name         string
	age          int
	favoritefood string
	vegetarian   bool
}

func main() {

	cliente1 := cliente{
		name:         "julio",
		age:          26,
		favoritefood: "temaki",
		vegetarian:   true,
	}
	cliente2 := cliente{
		name:         "carol",
		age:          23,
		favoritefood: "shimeji",
		vegetarian:   false,
	}
	cliente3 := cliente{
		name:         "paula",
		age:          34,
		favoritefood: "guioza",
		vegetarian:   false,
	}

	fmt.Println(cliente1)
	fmt.Println(cliente2)
	fmt.Println(cliente3)

}

package main

// built-in structs

import "fmt"

type person struct {
	name     string
	lastname string
	age      int
}
type collaborator struct {
	person
	office string
	salary float64
}

func main() {

	person1 := collaborator{
		person: person{
			name:     "julio",
			lastname: "scarceli",
			age:      26,
		},
		office: "developer",
		salary: 1500,
	}
	person2 := person{
		name:     "charlie",
		lastname: "sheen",
		age:      50,
	}
	person3 := person{"allan", "harper", 46}
	person4 := collaborator{person{"jake", "harper", 13}, "estagiario", 500}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)
	fmt.Println(person4)

	fmt.Println(person1.office)
	fmt.Println(person2.lastname)
	fmt.Println(person3.name)
	fmt.Println(person4.office)

}

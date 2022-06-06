package main

import "fmt"

// For variables with a wider scope, package level scope, we use var
var x int
var y string
var z bool
var i float64

func main() {
	//Variables declared without a corresponding
	//initialization are zero-valued. For example, the zero value for an int is 0
	fmt.Printf("%v\t%T\n", x, x)
	fmt.Printf("%v\t%T\n", y, y)
	fmt.Printf("%v\t%T\n", z, z)
	fmt.Printf("%v\t%T\n", i, i)

	// The := syntax is a shortcut for declaring and initializing a variable,
	// used to create new variables, inside code blocks

	x := 10
	y := "Hello, go"
	z := true
	i := 10.2

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(i)

	// "=" to assign values to already existing variables
	x = 15
	fmt.Print(x, "\n")

	// You can declare multiple variables at once
	var b, c int = 5, 6
	fmt.Println(b, c)

	//raw string literals
	a := `hello, \nlearner\n welcome to the study go/t`
	fmt.Println(a)

	//interpreted string literals
	a = "hello,\tlearner\n welcome to the study go\t"
	fmt.Println(a)

}

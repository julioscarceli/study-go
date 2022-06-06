package main

import "fmt"

var x bool // bool ( true or false )

func main() {

	fmt.Println(x) //zero value
	x = true
	fmt.Println(x) // assigned value

	// bools as a result of relational operators
	x = (10 > 100)
	y := (20 == 100)
	z := (30 < 100)
	i := (40 != 40)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(i)

}

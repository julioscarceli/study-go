package main

import "fmt"

var x [5]int
var y [6]int

func main() {
	x[0] = 1
	x[1] = 10
	fmt.Println(x[0], x[1])
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T", y)
	fmt.Println(len(x)) // to know the size of the array use len

	//----------------------------------------------------------------------
	var a [5]int
	fmt.Println("\n\nelements:", a)

	//We can set a value at an index using the array[index] = value syntax,
	//and get a value with array[index].
	a[3] = 50

	fmt.Println("set:", a)
	fmt.Println("get:", a[3])
	fmt.Println("len:", len(a))

	//Use this syntax to declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

}

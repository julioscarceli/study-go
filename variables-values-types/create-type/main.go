package main

import "fmt"

type hotdog int

var w hotdog

var j int

// creating type
func main() {

	fmt.Printf("%v\n%T\n", w, w)
	w = 42
	fmt.Println(w)

	//conversion type
	j = int(w)
	fmt.Println(j)
	fmt.Printf("%v\n%T\n", j, j)

}

package main

import "fmt"

func main() {
	// slice can have as many elements as you want
	array := [5]int{10, 20, 30, 40, 50}
	slice := []int{10, 20, 30, 40, 50}

	fmt.Println(array)
	fmt.Println(slice)

	fmt.Println(slice[3])
	slice[3] = 100
	fmt.Println(slice[3])

	//slice I can change the size, array can't
	slice2 := append(slice, 60)
	fmt.Println(slice2)
	fmt.Println("len:", len(slice2))

	// function append
	umaslice := []int{1, 2, 3, 4}
	outraslice := []int{7, 8, 9, 10}

	umaslice = append(umaslice, 5, 6)

	fmt.Println("\n\n", umaslice)

	umaslice = append(umaslice, outraslice...)
	fmt.Println(umaslice)

}

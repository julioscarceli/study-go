package main

import "fmt"

func main() {

	//for range ,the range goes through each element and gives me index and value
	slice := []string{"apple", "grape", "orange"}
	for i, v := range slice {
		fmt.Println("No indice", i, "temos o valor:", v)
	}

	// example
	slice1 := []int{10, 20, 30, 40, 50}
	total := 0

	for _, v := range slice1 {
		total += v // mesma coisa que total = total + valor

	}
	fmt.Println("\no valor total é de:", total)

	// example slicing slices
	color := []string{"blue", "green", "black", "yellow", "purple"}
	fatia := color[1:3]
	fmt.Println("\n", fatia)

	//example deleting slices
	numbers := []string{"twelve", "eleven", "five", "twenty", "one"}
	numbers = append(numbers[:2], numbers[4:]...)
	fmt.Println("\n", numbers)

}

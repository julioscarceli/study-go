package main

import "fmt"

// Crie e utilize uma função anônima.

func main() {
	x := []int{1, 1, 1, 1, 1}
	func(x []int) {
		total := 0
		for _, v := range x {
			total += v
		}
		fmt.Println(total)

	}(x)
}

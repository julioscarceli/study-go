package main

import "fmt"

func main() {
	ninjalevel := map[int]string{
		123: "beginner",
		456: "level2",
		789: "level3",
	}
	fmt.Println(ninjalevel)

	for k, v := range ninjalevel {
		fmt.Println(k, v)
	}

	delete(ninjalevel, 123) //deleting map elements
	fmt.Println("\n", ninjalevel)
}

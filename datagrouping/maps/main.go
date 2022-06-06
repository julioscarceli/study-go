package main

//Set key/value pairs using typical name[key] = val syntax
//Printing a map with e.g. fmt.Println will show all of its key/value pairs. ->

import "fmt"

func main() {

	friends := map[string]int{
		"carol":  20842856,
		"julio":  940732398,
		"gilson": 94777969,
	}
	fmt.Println(friends)
	fmt.Println(friends["carol"])

	friends["anderson"] = 48980999 //adding values to map
	fmt.Println(friends)
	fmt.Println(friends["anderson"], " \n ")

	//Sometimes you need to distinguish a missing entry from a zero value
	//comma ok idiom ->

	fmt.Println(friends["romario"])
	sera := friends["romario"]
	fmt.Println(sera)

	if sera, ok := friends["romario"]; !ok {
		fmt.Println("it does not have")
	} else {
		fmt.Println(sera)
	}
}

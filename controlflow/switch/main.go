package main

import "fmt"

func main() {

	x := 20
	switch {
	case x == 10:
		fmt.Println("x e igual a dez")
	case x < 10:
		fmt.Println("x e menor que dez")
	case x > 10:
		fmt.Println("x e maior que dez")
	}

	// example fallthrough and default in switch case
	whosthecompanytoday := "ninguem"
	switch whosthecompanytoday {
	case "julio":
		fmt.Println("who is in the company is julio")
	case "Jhon":
		fmt.Println("who is in the company is Jhon")
	case "Carol":
		fmt.Println("who is in the company is carol")
		fallthrough
	case "Sandro":
		fmt.Println("who is in the company is Sandro")
	default:
		fmt.Println("Nobody is in the company today")

	}

	//cases using two cases
	y := 5
	switch {
	case (y == 8), (y == 10):
		fmt.Println("the result is 8 or 10")
	case (y < 6):
		fmt.Println("3 ou 4")

	}

	//one more example
	favoriteseason := "Winter"
	switch {
	case favoriteseason == "Winter":
		fmt.Println("will travel between March and June")
	case favoriteseason == "summer":
		fmt.Println("will travel between december and march")
	case favoriteseason == "autumn":
		fmt.Println("will travel between June and september")
	}

}
